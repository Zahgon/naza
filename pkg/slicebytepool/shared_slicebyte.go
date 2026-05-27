// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package slicebytepool

import (
	"github.com/q191201771/naza/pkg/nazaatomic"
)

type SharedSliceByte struct {
	Core  []byte
	pool  SliceBytePool
	count nazaatomic.Uint32
}

type SharedSliceByteOption struct {
	pool SliceBytePool
}

var defaultSharedSliceByteOption = SharedSliceByteOption{
	pool: defaultPool,
}

type ModSharedSliceByteOption func(option *SharedSliceByteOption)

func WithPool(pool SliceBytePool) ModSharedSliceByteOption {
	_ = "STUB: not implemented"
	return *new(ModSharedSliceByteOption)
}

func NewSharedSliceByte(size int, modOptions ...ModSharedSliceByteOption) *SharedSliceByte {
	_ = "STUB: not implemented"
	return nil
}

func WrapSharedSliceByte(b []byte, modOptions ...ModSharedSliceByteOption) *SharedSliceByte {
	_ = "STUB: not implemented"
	return nil
}

func (ssb *SharedSliceByte) Ref() *SharedSliceByte { _ = "STUB: not implemented"; return nil }

func (ssb *SharedSliceByte) ReleaseIfNeeded() { _ = "STUB: not implemented"; return }
