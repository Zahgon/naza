// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package ratelimit

import (
	"errors"
	"sync"
	"time"

	"github.com/q191201771/naza/pkg/nazaatomic"
)

var ErrTokenNotEnough = errors.New("naza.ratelimit: token not enough")

// 令牌桶
type TokenBucket struct {
	capacity                  int
	prodTokenInterval         time.Duration
	prodTokenNumEveryInterval int

	disposeFlag nazaatomic.Bool

	mu        sync.Mutex
	available int
	cond      *sync.Cond
}

// @param capacity: 桶容量大小
// @param prodTokenIntervalMs: 生产令牌的时间间隔，单位毫秒
// @param prodTokenNumEveryInterval: 每次生产多少个令牌
func NewTokenBucket(capacity int, prodTokenIntervalMs int, prodTokenNumEveryInterval int) *TokenBucket {
	_ = "STUB: not implemented"
	return nil
}

func (tb *TokenBucket) TryAquire() error { _ = "STUB: not implemented"; return nil }

func (tb *TokenBucket) WaitUntilAquire() { _ = "STUB: not implemented"; return }

// 尝试获取相应数量的令牌，获取成功返回nil，获取失败返回ErrTokenNotEnough
// 如果获取失败，上层可自由选择多久后重试或丢弃本次任务
func (tb *TokenBucket) TryAquireWithNum(num int) error { _ = "STUB: not implemented"; return nil }

// 阻塞直到获取到相应数量的令牌
func (tb *TokenBucket) WaitUntilAquireWithNum(num int) { _ = "STUB: not implemented"; return }

// 等待下次令牌生产时被唤醒
// wait的内部会将自身添加到事件监听队列中然后释放锁，当接收到事件时，内部会重新获取锁然后返回

// 销毁令牌桶
func (tb *TokenBucket) Dispose() { _ = "STUB: not implemented"; return }

func (tb *TokenBucket) asyncProdToken() { _ = "STUB: not implemented"; return }

// It is allowed but not required for the caller to hold c.L
// during the call.

func (tb *TokenBucket) checkAquireNum(num int) { _ = "STUB: not implemented"; return }
