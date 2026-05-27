// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package snowflake

import (
	"errors"
	"sync"
)

var (
	ErrInitial = errors.New("lal.snowflake: initial error")
	ErrGen     = errors.New("lal.snowflake: gen error")
)

type Option struct {
	DataCenterIdBits int   // 数据中心编号字段在所生成 ID 所占的位数，取值范围见 validate 函数
	WorkerIdBits     int   // 节点编号
	SequenceBits     int   // 递增序列
	Twepoch          int64 // 基准时间点
	AlwaysPositive   bool  // 是否只生成正数 ID，如果是，则时间戳所占位数会减少1位
}

var defaultOption = Option{
	DataCenterIdBits: 5,
	WorkerIdBits:     5,
	SequenceBits:     12,
	Twepoch:          int64(1288834974657), // 对应现实时间： 2010/11/4 9:42:54.657
	AlwaysPositive:   false,
}

type Node struct {
	dataCenterId int64
	workerId     int64
	option       Option

	seqMask           uint32
	workerIdShift     uint32
	dataCenterIdShift uint32
	timestampShift    uint32

	mu     sync.Mutex
	lastTs int64
	seq    uint32
}

type ModOption func(option *Option)

// dataCenterId 和 workerId 的取值范围取决于 DataCenterIdBits 和 WorkerIdBits
// 假设 DataCenterIdBits 为 5，则 dataCenterId 取值范围为 [0, 32]
func New(dataCenterId int, workerId int, modOptions ...ModOption) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *Node) Gen(nowUnixMs ...int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// 当前 Unix 时间戳可由外部传入

// 时间戳回退，返回错误

// 时间戳相同时，使用递增序号解决冲突

// 递增序号翻转为 0，表示该时间戳下的序号已经全部用完，阻塞等待系统时间增长

// 如果保证只返回正数，则生成的 ID 的最高位，也即时间戳的最高位保持为 0

// 用所有字段组合生成 ID 返回

func validate(dataCenterId int, workerId int, option Option) error {
	_ = "STUB: not implemented"
	return nil
}

// 位的数量对应的最大值，该函数也可以叫做 bitsToMask
func bitsToMax(bits int) int {
	_ = "STUB: not implemented"
	// -1 表示所有位都为 1
	return 0
}

// 将 <num> 的第 <index> 设置为 0
func clearBit(num int64, index uint32) int64 { _ = "STUB: not implemented"; return 0 }
