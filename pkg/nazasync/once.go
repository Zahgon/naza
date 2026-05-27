// Copyright 2022, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazasync

import (
	"sync"
)

type StdOnce struct {
	core sync.Once
}

func (o *StdOnce) Do(f func()) { _ = "STUB: not implemented"; return }

type NonblockingOnce struct {
	done uint32
}

func (o *NonblockingOnce) Do(f func()) { _ = "STUB: not implemented"; return }
