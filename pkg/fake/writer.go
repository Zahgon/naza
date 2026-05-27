// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package fake

import (
	"bytes"
	"errors"
)

type WriterType uint8

const (
	WriterTypeDoNothing WriterType = iota
	WriterTypeReturnError
	WriterTypeIntoBuffer
)

var (
	ErrFakeWriter = errors.New("naza.fake: a fake writer error")
)

type Writer struct {
	t     WriterType
	ts    map[uint32]WriterType
	count uint32
	B     bytes.Buffer
}

func NewWriter(t WriterType) *Writer { _ = "STUB: not implemented"; return nil }

// 为某些写操作指定特定的类型，次数从 0 开始计数
func (w *Writer) SetSpecificType(ts map[uint32]WriterType) { _ = "STUB: not implemented"; return }

func (w *Writer) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
