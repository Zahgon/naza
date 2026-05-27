// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package taskpool

import (
	"sync"
)

type taskWrapper struct {
	taskFn      TaskFn
	param       []interface{}
	disposeFlag bool
}

type pool struct {
	maxWorkerNum int

	m sync.Mutex
	//totalWorkerNum int
	idleWorkerList []*worker
	blockTaskList  []taskWrapper
	allWorkerList  []*worker
	disposeFlag    bool
}

func newPool(option Option) *pool { _ = "STUB: not implemented"; return nil }

func (p *pool) Go(task TaskFn, param ...interface{}) { _ = "STUB: not implemented"; return }

// 还有空闲worker

// 无空闲worker

// 无最大worker限制，或还未达到限制

// 已达到限制

func (p *pool) KillIdleWorkers() { _ = "STUB: not implemented"; return }

func (p *pool) Dispose(t DisposeType) { _ = "STUB: not implemented"; return }

// noop

func (p *pool) GetCurrentStatus() Status { _ = "STUB: not implemented"; return *new(Status) }

func (p *pool) newWorker() *worker { _ = "STUB: not implemented"; return nil }

func (p *pool) newWorkerWithTask(task taskWrapper) { _ = "STUB: not implemented"; return }

func (p *pool) onIdle(w *worker) { _ = "STUB: not implemented"; return }

// 没有等待执行的任务

func (p *pool) onDispose(w *worker) { _ = "STUB: not implemented"; return }
