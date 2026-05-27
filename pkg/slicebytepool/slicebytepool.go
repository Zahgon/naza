// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package slicebytepool

import "github.com/q191201771/naza/pkg/nazaatomic"

var (
	minSize = 1024
	maxSize = 1073741824
)

type sliceBytePool struct {
	strategy        Strategy
	capToFreeBucket map[int]Bucket
	status          statusAtomic
}

type statusAtomic struct {
	getCount  nazaatomic.Int64
	putCount  nazaatomic.Int64
	hitCount  nazaatomic.Int64
	sizeBytes nazaatomic.Int64
}

func (bp *sliceBytePool) Get(size int) []byte { _ = "STUB: not implemented"; return nil }

func (bp *sliceBytePool) Put(buf []byte) { _ = "STUB: not implemented"; return }

func (bp *sliceBytePool) RetrieveStatus() Status { _ = "STUB: not implemented"; return *new(Status) }

// @return 范围为 [2, 4, 8, 16, ..., 1073741824]，如果大于等于1073741824，则直接返回n
func up2power(n int) int { _ = "STUB: not implemented"; return 0 }

// @return 范围为 [2, 4, 8, 16, ..., 1073741824]
func down2power(n int) int { _ = "STUB: not implemented"; return 0 }
