// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package retry

import (
	"time"

	"github.com/luci/luci-go/common/clock"
	"golang.org/x/net/context"
)

// Limited is an Iterator implementation that may be limited by a maximum number
// of retries and/or time.
type Limited struct {
	// Delay is the next generated delay.
	Delay time.Duration

	// Retries, if >= 0, is the number of remaining retries. If <0, no retry
	// count will be applied.
	Retries int

	// MaxTotal is the maximum total elapsed time. If <= 0, no maximum will be
	// enforced.
	MaxTotal time.Duration

	// The time when the generator initially started.
	startTime time.Time
}

var _ Iterator = (*Limited)(nil)

// Next implements the Iterator interface.
func (i *Limited) Next(ctx context.Context, _ error) time.Duration {
	switch {
	case i.Retries == 0:
		return Stop
	case i.Retries > 0:
		i.Retries--
	}

	// If there is a maximum total time, enforce it.
	if i.MaxTotal > 0 {
		now := clock.Now(ctx)
		if i.startTime.IsZero() {
			i.startTime = now
		}

		var elapsed time.Duration
		if now.After(i.startTime) {
			elapsed = now.Sub(i.startTime)
		}

		// Remaining time is the difference between total allowed time and elapsed
		// time.
		remaining := i.MaxTotal - elapsed
		if remaining <= 0 {
			// No more time!
			i.Retries = 0
			return Stop
		}
	}

	return i.Delay
}
