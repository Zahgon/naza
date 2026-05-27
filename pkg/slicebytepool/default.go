// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package slicebytepool

var defaultPool SliceBytePool

func Get(size int) []byte { _ = "STUB: not implemented"; return nil }

func Put(buf []byte) { _ = "STUB: not implemented"; return }

func RetrieveStatus() Status { _ = "STUB: not implemented"; return *new(Status) }

func Init(strategy Strategy) { _ = "STUB: not implemented"; return }

func init() {
	Init(StrategyMultiSlicePoolBucket)
}
