// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

// Package bitrate 平滑计算比特率（码率）
package bitrate

import (
	"sync"
)

type Bitrate interface {
	// Add
	//
	// @param nowUnixMs: 变参，可选择从外部传入当前 unix 时间戳，单位毫秒
	//
	Add(bytes int, nowUnixMs ...int64)

	Rate(nowUnixMs ...int64) float32
}

type Unit uint8

const (
	UnitBitPerSec Unit = iota + 1
	UnitBytePerSec
	UnitKbitPerSec
	UnitKbytePerSec
)

// TODO chef: 考虑支持配置是否在内部使用锁
type Option struct {
	WindowMs int
	Unit     Unit
}

var defaultOption = Option{
	WindowMs: 1000,
	Unit:     UnitKbitPerSec,
}

type ModOption func(option *Option)

func New(modOptions ...ModOption) Bitrate { _ = "STUB: not implemented"; return *new(Bitrate) }

type bitrate struct {
	option Option

	mu          sync.Mutex
	bucketSlice []bucket
}

type bucket struct {
	n int
	t int64 // unix 时间戳，单位毫秒
}

func (b *bitrate) Add(bytes int, nowUnixMs ...int64) { _ = "STUB: not implemented"; return }

func (b *bitrate) Rate(nowUnixMs ...int64) float32 { _ = "STUB: not implemented"; return 0 }

func (b *bitrate) sweepStale(now int64) { _ = "STUB: not implemented"; return }
