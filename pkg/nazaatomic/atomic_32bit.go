// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

//go:build 386 || arm || mips || mipsle
// +build 386 arm mips mipsle

package nazaatomic

import (
	"sync"
)

// 注意，因为32位系统下，使用标准库中的atomic操作64位整型有bug，所以32位系统原子操作64位整型时，我们使用mutex

type Int64 struct {
	mu   sync.Mutex
	core int64
}

type Uint64 struct {
	mu   sync.Mutex
	core uint64
}

// ----------------------------------------------------------------------------

func (obj *Uint64) Load() uint64 { _ = "STUB: not implemented"; return 0 }

func (obj *Uint64) Store(val uint64) { _ = "STUB: not implemented"; return }

func (obj *Uint64) Add(delta uint64) (new uint64) { _ = "STUB: not implemented"; return 0 }

// @param delta 举例，传入3，则减3
func (obj *Uint64) Sub(delta uint64) (new uint64) { _ = "STUB: not implemented"; return 0 }

func (obj *Uint64) Increment() (new uint64) { _ = "STUB: not implemented"; return 0 }

func (obj *Uint64) Decrement() (new uint64) { _ = "STUB: not implemented"; return 0 }

func (obj *Uint64) CompareAndSwap(old uint64, new uint64) (swapped bool) {
	_ = "STUB: not implemented"
	return false
}

func (obj *Uint64) Swap(new uint64) (old uint64) { _ = "STUB: not implemented"; return 0 }

// ----------------------------------------------------------------------------

func (obj *Int64) Load() int64 { _ = "STUB: not implemented"; return 0 }

func (obj *Int64) Store(val int64) { _ = "STUB: not implemented"; return }

func (obj *Int64) Add(delta int64) (new int64) { _ = "STUB: not implemented"; return 0 }

// @param delta 举例，传入3，则减3
func (obj *Int64) Sub(delta int64) (new int64) { _ = "STUB: not implemented"; return 0 }

func (obj *Int64) Increment() (new int64) { _ = "STUB: not implemented"; return 0 }

func (obj *Int64) Decrement() (new int64) { _ = "STUB: not implemented"; return 0 }

func (obj *Int64) CompareAndSwap(old int64, new int64) (swapped bool) {
	_ = "STUB: not implemented"
	return false
}

func (obj *Int64) Swap(new int64) (old int64) { _ = "STUB: not implemented"; return 0 }
