// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package slicebytepool

import (
	"sync"
)

type SliceBucket struct {
	m    sync.Mutex
	core [][]byte
}

func NewSliceBucket() *SliceBucket { _ = "STUB: not implemented"; return nil }

func (b *SliceBucket) Get(size int) []byte { _ = "STUB: not implemented"; return nil }

func (b *SliceBucket) Put(buf []byte) { _ = "STUB: not implemented"; return }
