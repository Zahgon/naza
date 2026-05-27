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

type StdPoolBucket struct {
	core *sync.Pool
}

func NewStdPoolBucket() *StdPoolBucket { _ = "STUB: not implemented"; return nil }

func (b *StdPoolBucket) Get(size int) []byte { _ = "STUB: not implemented"; return nil }

func (b *StdPoolBucket) Put(buf []byte) { _ = "STUB: not implemented"; return }
