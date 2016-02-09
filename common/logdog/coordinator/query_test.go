// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package coordinator

import (
	"errors"
	"fmt"
	"testing"

	"github.com/luci/luci-go/common/api/logdog_coordinator/logs/v1"
	"github.com/luci/luci-go/common/clock/testclock"
	"github.com/luci/luci-go/common/grpcutil"
	"github.com/luci/luci-go/common/proto/google"
	"github.com/luci/luci-go/common/proto/logdog/logpb"
	"github.com/luci/luci-go/common/testing/prpctest"
	"golang.org/x/net/context"

	. "github.com/smartystreets/goconvey/convey"
)

type testQueryLogsService struct {
	testLogsServiceBase

	LR logs.QueryRequest
	H  func(*logs.QueryRequest) (*logs.QueryResponse, error)
}

func (s *testQueryLogsService) Query(c context.Context, req *logs.QueryRequest) (*logs.QueryResponse, error) {
	s.LR = *req
	if h := s.H; h != nil {
		return s.H(req)
	}
	return nil, errors.New("not implemented")
}

func gen(name string, state *logs.LogStreamState) *logs.QueryResponse_Stream {
	return &logs.QueryResponse_Stream{
		Path:  fmt.Sprintf("test/+/%s", name),
		State: state,
		Desc: &logpb.LogStreamDescriptor{
			Prefix: "test",
			Name:   name,
		},
	}
}

func shouldHaveLogStreams(actual interface{}, expected ...interface{}) string {
	a := actual.([]*LogStream)

	aList := make([]string, len(a))
	for i, ls := range a {
		aList[i] = string(ls.Path)
	}

	eList := make([]string, len(expected))
	for i, exp := range expected {
		eList[i] = exp.(string)
	}

	return ShouldResemble(aList, eList)
}

func TestClientQuery(t *testing.T) {
	t.Parallel()

	Convey(`A testing Client`, t, func() {
		now := testclock.TestTimeLocal
		c := context.Background()

		ts := prpctest.Server{}
		svc := testQueryLogsService{}
		logs.RegisterLogsServer(&ts, &svc)

		// Create a testing server and client.
		ts.Start(c)
		defer ts.Close()

		prpcClient, err := ts.NewClient()
		if err != nil {
			panic(err)
		}
		client := Client{
			C: logs.NewLogsPRPCClient(prpcClient),
		}

		Convey(`When making a query request`, func() {
			q := Query{
				Path: "**/+/**",
				Tags: map[string]string{
					"foo": "bar",
					"baz": "qux",
				},
				ContentType: "application/text",
				Before:      now,
				After:       now,
				Terminated:  Yes,
				Archived:    No,
				Purged:      Both,
				State:       true,
			}

			var results []*LogStream
			accumulate := func(s *LogStream) bool {
				results = append(results, s)
				return true
			}

			Convey(`Can accumulate results across queries.`, func() {
				// This handler will return a single query per request, as well as a
				// non-empty Next pointer for the next query element. It progresses
				// "a" => "b" => "final" => "".
				svc.H = func(req *logs.QueryRequest) (*logs.QueryResponse, error) {
					r := logs.QueryResponse{}
					switch req.Next {
					case "":
						r.Streams = append(r.Streams, gen("a", nil))
						r.Next = "b"
					case "b":
						r.Streams = append(r.Streams, gen("b", nil))
						r.Next = "final"
					case "final":
						r.Streams = append(r.Streams, gen("final", nil))
					default:
						return nil, errors.New("invalid cursor")
					}
					return &r, nil
				}

				So(client.Query(c, &q, accumulate), ShouldBeNil)
				So(results, shouldHaveLogStreams, "test/+/a", "test/+/b", "test/+/final")
			})

			Convey(`Will stop invoking the callback if it returns false.`, func() {
				// This handler will return three query results, "a", "b", and "c".
				svc.H = func(*logs.QueryRequest) (*logs.QueryResponse, error) {
					return &logs.QueryResponse{
						Streams: []*logs.QueryResponse_Stream{
							gen("a", nil),
							gen("b", nil),
							gen("c", nil),
						},
						Next: "infiniteloop",
					}, nil
				}

				accumulate = func(s *LogStream) bool {
					results = append(results, s)
					return len(results) < 3
				}
				So(client.Query(c, &q, accumulate), ShouldBeNil)
				So(results, shouldHaveLogStreams, "test/+/a", "test/+/b", "test/+/c")
			})

			Convey(`Will properly handle state and protobuf deserialization.`, func() {
				svc.H = func(*logs.QueryRequest) (*logs.QueryResponse, error) {
					return &logs.QueryResponse{
						Streams: []*logs.QueryResponse_Stream{
							gen("a", &logs.LogStreamState{
								Created: google.NewTimestamp(now),
								Updated: google.NewTimestamp(now),
							}),
						},
					}, nil
				}

				So(client.Query(c, &q, accumulate), ShouldBeNil)
				So(results, shouldHaveLogStreams, "test/+/a")
				So(results[0], ShouldResemble, &LogStream{
					Path: "test/+/a",
					Desc: &logpb.LogStreamDescriptor{Prefix: "test", Name: "a"},
					State: &StreamState{
						Created: now.UTC(),
						Updated: now.UTC(),
					},
				})
			})

			Convey(`Can query for stream types`, func() {
				svc.H = func(*logs.QueryRequest) (*logs.QueryResponse, error) {
					return &logs.QueryResponse{}, nil
				}

				Convey(`Text`, func() {
					q.StreamType = Text
					So(client.Query(c, &q, accumulate), ShouldBeNil)
					So(svc.LR.StreamType, ShouldResemble, &logs.QueryRequest_StreamTypeFilter{Value: logpb.StreamType_TEXT})
				})

				Convey(`Binary`, func() {
					q.StreamType = Binary
					So(client.Query(c, &q, accumulate), ShouldBeNil)
					So(svc.LR.StreamType, ShouldResemble, &logs.QueryRequest_StreamTypeFilter{Value: logpb.StreamType_BINARY})
				})

				Convey(`Datagram`, func() {
					q.StreamType = Datagram
					So(client.Query(c, &q, accumulate), ShouldBeNil)
					So(svc.LR.StreamType, ShouldResemble, &logs.QueryRequest_StreamTypeFilter{Value: logpb.StreamType_DATAGRAM})
				})
			})

			Convey(`Will return ErrNoAccess if unauthenticated.`, func() {
				svc.H = func(*logs.QueryRequest) (*logs.QueryResponse, error) {
					return nil, grpcutil.Unauthenticated
				}

				So(client.Query(c, &q, accumulate), ShouldEqual, ErrNoAccess)
			})

			Convey(`Will return ErrNoAccess if permission denied.`, func() {
				svc.H = func(*logs.QueryRequest) (*logs.QueryResponse, error) {
					return nil, grpcutil.Unauthenticated
				}

				So(client.Query(c, &q, accumulate), ShouldEqual, ErrNoAccess)
			})
		})
	})
}
