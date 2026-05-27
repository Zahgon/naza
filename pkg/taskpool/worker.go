// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package taskpool

type worker struct {
	taskChan chan taskWrapper
	p        *pool
}

func NewWorker(p *pool) *worker { _ = "STUB: not implemented"; return nil }

func (w *worker) Start() { _ = "STUB: not implemented"; return }

func (w *worker) Stop() { _ = "STUB: not implemented"; return }

func (w *worker) Go(t taskWrapper) { _ = "STUB: not implemented"; return }
