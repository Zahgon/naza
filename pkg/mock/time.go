// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package mock

import (
	"sync"
	"time"
)

// TODO(chef): [feat] 增加Clock::NewTicker

type Clock interface {
	// Now NewTimer ...
	//
	// 标准库中的操作集合, stdClock和mockClock都有对应的实现
	//
	Now() time.Time
	NewTimer(d time.Duration) *Timer
	Sleep(d time.Duration)

	// Add Set ...
	//
	// mockClock使用以下这些函数来修改当前时间
	// 注意，如果是stdClock，则没有必要调用以下函数（调用以下函数为空实现）
	//
	Add(d time.Duration)
	Set(t time.Time)
}

func NewStdClock() Clock { _ = "STUB: not implemented"; return *new(Clock) }

func NewFakeClock() Clock {
	_ = "STUB: not implemented"
	return *

	// ---------------------------------------------------------------------------------------------------------------------
	new(Clock)
}

type stdClock struct {
}

func (c *stdClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *stdClock) NewTimer(d time.Duration) *Timer { _ = "STUB: not implemented"; return nil }

func (c *stdClock) Sleep(d time.Duration) { _ = "STUB: not implemented"; return }

func (c *stdClock) Add(d time.Duration) {
	_ = "STUB: not implemented"
	// noop
	return
}

func (c *stdClock) Set(t time.Time) {
	_ = "STUB: not implemented"
	// noop

	// ---------------------------------------------------------------------------------------------------------------------
	return
}

type fakeClock struct {
	mu     sync.Mutex
	now    time.Time
	timers timers
}

func (fc *fakeClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (fc *fakeClock) NewTimer(d time.Duration) *Timer { _ = "STUB: not implemented"; return nil }

func (fc *fakeClock) Sleep(d time.Duration) {
	_ = "STUB: not implemented"
	// TODO(chef): [feat] 实现和add、set挂钩的Sleep，内部用Timer实现等待
	// 当前的使用场景都是测试场景，直接快速跳过Sleep以及能够满足需求
	return
}

func (fc *fakeClock) Add(d time.Duration) { _ = "STUB: not implemented"; return }

func (fc *fakeClock) Set(t time.Time) { _ = "STUB: not implemented"; return }

func (fc *fakeClock) resetTimerWithLock(t *Timer, d time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

func (fc *fakeClock) stopTimerWithLock(t *Timer) bool { _ = "STUB: not implemented"; return false }

func (fc *fakeClock) addTimer(t *Timer) { _ = "STUB: not implemented"; return }

func (fc *fakeClock) delTimer(t *Timer) { _ = "STUB: not implemented"; return }

func (fc *fakeClock) ringTimersIfNeeded() { _ = "STUB: not implemented"; return }

// TODO(chef): [perf] 用有序map
type timers []*Timer

func (t timers) Len() int { _ = "STUB: not implemented"; return 0 }

func (t timers) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (t timers) Swap(i, j int) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------------------------------------------

type Timer struct {
	C <-chan time.Time

	stdTimer *time.Timer

	fc      *fakeClock
	timing  time.Time
	c       chan time.Time
	expired bool
	stopped bool
}

func (t *Timer) Reset(d time.Duration) bool { _ = "STUB: not implemented"; return false }

func (t *Timer) Stop() bool { _ = "STUB: not implemented"; return false }
