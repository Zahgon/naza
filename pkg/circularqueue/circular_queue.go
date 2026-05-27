// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package circularqueue

import "errors"

// 底层基于切片实现的固定容量大小的FIFO的环形队列

var ErrCircularQueue = errors.New("circular queue: fxxk")

type CircularQueue struct {
	capacity int
	core     []interface{}
	first    int
	last     int
}

func New(capacity int) *CircularQueue { _ = "STUB: not implemented"; return nil }

// @return 如果队列满了，则返回错误
func (c *CircularQueue) PushBack(v interface{}) error { _ = "STUB: not implemented"; return nil }

// @return 如果队列为空，则返回错误
func (c *CircularQueue) PopFront() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// @return 如果队列为空，则返回错误
func (c *CircularQueue) Front() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// @return 如果队列为空，则返回错误
func (c *CircularQueue) Back() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// 获取第i个元素
func (c *CircularQueue) At(i int) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *CircularQueue) Size() int { _ = "STUB: not implemented"; return 0 }

func (c *CircularQueue) Full() bool { _ = "STUB: not implemented"; return false }

func (c *CircularQueue) Empty() bool { _ = "STUB: not implemented"; return false }
