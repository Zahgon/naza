// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazasync

import (
	"sync"
	"time"

	"github.com/q191201771/naza/pkg/unique"
)

// 用于debug锁方面的问题

var uniqueGen *unique.SingleGenerator

type Mutex struct {
	core          sync.Mutex
	startHoldTime time.Time

	genUniqueKeyOnce sync.Once
	uniqueKey        string
}

func (m *Mutex) Lock() { _ = "STUB: not implemented"; return }

func (m *Mutex) Unlock() { _ = "STUB: not implemented"; return }

// 运行时自己会检查对没有加锁的mutex进行Unlock调用的情况

var globalMutexManager = NewMutexManager()

// 注意，key是由mutex唯一ID加上协程ID组合而成
type MutexManager struct {
	mu                   sync.Mutex
	waitAcquireContainer map[string]time.Time
	holdContainer        map[string]time.Time
}

func NewMutexManager() *MutexManager { _ = "STUB: not implemented"; return nil }

func (m *MutexManager) printTmpDebug() { _ = "STUB: not implemented"; return }

func (m *MutexManager) beforeAcquireLock(uk string, gid int64) { _ = "STUB: not implemented"; return }

// 当前协程已持有锁，再次重入

func (m *MutexManager) afterAcquireLock(uk string, gid int64) { _ = "STUB: not implemented"; return }

func (m *MutexManager) afterUnlock(uk string, gid int64) { _ = "STUB: not implemented"; return }

// 有可能是a协程Lock，b协程Unlock

func init() {
	uniqueGen = unique.NewSingleGenerator("MUTEX")
}
