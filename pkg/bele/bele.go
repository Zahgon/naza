// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

// Package bele 提供了大小端的转换操作
//
// be是big endian的缩写，即大端
// le是little endian的缩写，即小端
//
// assume local is `le`
package bele

import (
	"io"
)

// ----- 反序列化 -----

func BeUint16(p []byte) uint16 { _ = "STUB: not implemented"; return 0 }

func BeUint24(p []byte) uint32 { _ = "STUB: not implemented"; return 0 }

func BeUint32(p []byte) (ret uint32) { _ = "STUB: not implemented"; return 0 }

func BeUint64(p []byte) (ret uint64) { _ = "STUB: not implemented"; return 0 }

func BeFloat64(p []byte) (ret float64) { _ = "STUB: not implemented"; return 0 }

func LeUint32(p []byte) (ret uint32) { _ = "STUB: not implemented"; return 0 }

func LeUint16(p []byte) (ret uint16) { _ = "STUB: not implemented"; return 0 }

func ReadBytes(r io.Reader, n int) ([]byte, error) {
	_ = "STUB: not implemented"

	// 原生Read函数，读不够时，会在第一次调用时读入剩余的数据，并返回读入数据的真实长度，以及nil值的error
	// 在下一次Read时，才返回EOF
	// 这里我们在第一次读不够时，就直接返回EOF。（但是也会把剩余的数据读取到）
	return nil, nil
}

func ReadString(r io.Reader, n int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ReadUint8(r io.Reader) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadBeUint16(r io.Reader) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadBeUint24(r io.Reader) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadBeUint32(r io.Reader) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadBeUint64(r io.Reader) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadLeUint32(r io.Reader) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadLeUint16(r io.Reader) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

// ----- 序列化 -----

func BePutUint16(out []byte, in uint16) { _ = "STUB: not implemented"; return }

func BePutUint24(out []byte, in uint32) { _ = "STUB: not implemented"; return }

func BePutUint32(out []byte, in uint32) { _ = "STUB: not implemented"; return }

func BePutUint64(out []byte, in uint64) { _ = "STUB: not implemented"; return }

func LePutUint32(out []byte, in uint32) { _ = "STUB: not implemented"; return }

func LePutUint16(out []byte, in uint16) { _ = "STUB: not implemented"; return }

func WriteBeUint24(writer io.Writer, in uint32) error { _ = "STUB: not implemented"; return nil }

func WriteBe(writer io.Writer, in interface{}) error { _ = "STUB: not implemented"; return nil }

func WriteLe(writer io.Writer, in interface{}) error { _ = "STUB: not implemented"; return nil }
