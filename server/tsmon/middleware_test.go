// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tsmon

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/net/context"

	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/tsmon"
	"go.chromium.org/luci/common/tsmon/field"
	"go.chromium.org/luci/common/tsmon/monitor"
	"go.chromium.org/luci/common/tsmon/store"
	"go.chromium.org/luci/common/tsmon/store/storetest"
	"go.chromium.org/luci/common/tsmon/target"
	"go.chromium.org/luci/common/tsmon/types"
	"go.chromium.org/luci/server/router"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMiddleware(t *testing.T) {
	t.Parallel()

	metric := &storetest.FakeMetric{
		types.MetricInfo{"m", "", []field.Field{}, types.CumulativeIntType},
		types.MetricMetadata{},
	}

	runMiddlware := func(c context.Context, state *State, cb func(*router.Context)) {
		rec := httptest.NewRecorder()
		router.RunMiddleware(
			&router.Context{Context: c, Writer: rec, Request: &http.Request{}},
			router.NewMiddlewareChain(state.Middleware),
			cb,
		)
		So(rec.Code, ShouldEqual, http.StatusOK)
	}

	incrMetric := func(c *router.Context) {
		So(store.IsNilStore(tsmon.Store(c.Context)), ShouldBeFalse)
		So(tsmon.Store(c.Context).Incr(c.Context, metric, time.Time{}, []interface{}{}, int64(1)), ShouldBeNil)
	}

	readMetric := func(c context.Context) interface{} {
		value, err := tsmon.Store(c).Get(c, metric, time.Time{}, []interface{}{})
		if err != nil {
			panic(err)
		}
		return value
	}

	Convey("With fakes", t, func() {
		c, clock := buildTestContext()
		state, monitor, allocator := buildTestState()
		tsmon.Register(c, metric)

		Convey("Pings TaskNumAllocator and waits for number", func() {
			So(len(allocator.taskIDs), ShouldEqual, 0)
			allocator.taskNum = -1 // no number yet

			// Do the flush.
			state.nextFlush = clock.Now()
			runMiddlware(c, state, incrMetric)

			// Called the allocator.
			So(len(allocator.taskIDs), ShouldEqual, 1)
			// Shouldn't flush since the instance entity doesn't have a task number yet.
			So(len(monitor.Cells), ShouldEqual, 0)

			// Wait until next expected flush.
			clock.Add(time.Minute)
			allocator.taskNum = 0 // got the number!

			// Do the flush again.
			state.nextFlush = clock.Now()
			runMiddlware(c, state, incrMetric)

			// Flushed stuff this time.
			So(len(monitor.Cells), ShouldEqual, 1)

			// The value should still be set.
			So(readMetric(c), ShouldEqual, int64(2))
		})

		Convey("Resets cumulative metrics", func() {
			// Do the flush.
			state.nextFlush = clock.Now()
			runMiddlware(c, state, incrMetric)

			// Flushed stuff.
			So(len(monitor.Cells), ShouldEqual, 1)
			monitor.Cells = nil

			// The value is set.
			So(readMetric(c), ShouldEqual, int64(1))

			// Lost the task number.
			allocator.taskNum = -1

			// Do the flush again.
			state.nextFlush = clock.Now()
			runMiddlware(c, state, incrMetric)

			// No stuff is sent, and cumulative metrics are reset.
			So(len(monitor.Cells), ShouldEqual, 0)
			So(readMetric(c), ShouldEqual, nil)
		})

		Convey("Dynamic enable and disable works", func() {
			// Enabled. Store is not nil.
			runMiddlware(c, state, func(c *router.Context) {
				So(store.IsNilStore(tsmon.Store(c.Context)), ShouldBeFalse)
			})

			// Disabled. Store is nil.
			state.testingSettings.Enabled = false
			runMiddlware(c, state, func(c *router.Context) {
				So(store.IsNilStore(tsmon.Store(c.Context)), ShouldBeTrue)
			})
		})
	})
}

////////////////////////////////////////////////////////////////////////////////

func buildTestContext() (context.Context, testclock.TestClock) {
	c := context.Background()
	c, clock := testclock.UseTime(c, testclock.TestTimeUTC)
	c, _, _ = tsmon.WithFakes(c)
	return c, clock
}

func buildTestState() (*State, *monitor.Fake, *fakeNumAllocator) {
	mon := &monitor.Fake{}
	allocator := &fakeNumAllocator{}
	return &State{
		Target: func(c context.Context) target.Task {
			return target.Task{
				DataCenter:  "datacenter",
				ServiceName: "app-id",
				JobName:     "service-name",
				HostName:    "12345-version",
			}
		},
		TaskID:           func(c context.Context) string { return "some.task.id" },
		TaskNumAllocator: allocator,
		IsDevMode:        false, // hit same paths as prod
		testingMonitor:   mon,
		testingSettings: &tsmonSettings{
			Enabled:          true,
			FlushIntervalSec: 60,
		},
	}, mon, allocator
}

type fakeNumAllocator struct {
	taskNum int
	taskIDs []string
}

func (a *fakeNumAllocator) NotifyTaskIsAlive(c context.Context, taskID string) (int, error) {
	a.taskIDs = append(a.taskIDs, taskID)
	if a.taskNum == -1 {
		return 0, ErrNoTaskNumber
	}
	return a.taskNum, nil
}