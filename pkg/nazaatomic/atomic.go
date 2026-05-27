// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazaatomic

type Bool struct {
	core Int32
}

type Int32 struct {
	core int32
}

type Uint32 struct {
	core uint32
}

// ----------------------------------------------------------------------------

func (obj *Int32) Load() int32 { _ = "STUB: not implemented"; return 0 }

func (obj *Int32) Store(val int32) { _ = "STUB: not implemented"; return }

func (obj *Int32) Add(delta int32) (new int32) { _ = "STUB: not implemented"; return 0 }

// @param delta 举例，传入3，则减3
func (obj *Int32) Sub(delta int32) (new int32) { _ = "STUB: not implemented"; return 0 }

func (obj *Int32) Increment() (new int32) { _ = "STUB: not implemented"; return 0 }

func (obj *Int32) Decrement() (new int32) { _ = "STUB: not implemented"; return 0 }

func (obj *Int32) CompareAndSwap(old int32, new int32) (swapped bool) {
	_ = "STUB: not implemented"
	return false
}

func (obj *Int32) Swap(new int32) (old int32) { _ = "STUB: not implemented"; return 0 }

// ----------------------------------------------------------------------------

func (obj *Uint32) Load() uint32 { _ = "STUB: not implemented"; return 0 }

func (obj *Uint32) Store(val uint32) { _ = "STUB: not implemented"; return }

func (obj *Uint32) Add(delta uint32) (new uint32) { _ = "STUB: not implemented"; return 0 }

// @param delta 举例，传入3，则减3
func (obj *Uint32) Sub(delta uint32) (new uint32) { _ = "STUB: not implemented"; return 0 }

func (obj *Uint32) Increment() (new uint32) { _ = "STUB: not implemented"; return 0 }

func (obj *Uint32) Decrement() (new uint32) { _ = "STUB: not implemented"; return 0 }

func (obj *Uint32) CompareAndSwap(old uint32, new uint32) (swapped bool) {
	_ = "STUB: not implemented"
	return false
}

func (obj *Uint32) Swap(new uint32) (old uint32) { _ = "STUB: not implemented"; return 0 }

// ----------------------------------------------------------------------------

func (obj *Bool) Load() bool { _ = "STUB: not implemented"; return false }

func (obj *Bool) Store(val bool) { _ = "STUB: not implemented"; return }

func (obj *Bool) CompareAndSwap(old bool, new bool) (swapped bool) {
	_ = "STUB: not implemented"
	return false
}

func (obj *Bool) Swap(new bool) (old bool) { _ = "STUB: not implemented"; return false }

func booltoint32(val bool) int32 { _ = "STUB: not implemented"; return 0 }

func int32tobool(val int32) bool { _ = "STUB: not implemented"; return false }
