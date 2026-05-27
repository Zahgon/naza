// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazabits

import "errors"

var ErrNazaBits = errors.New("nazabits: fxxk")

// BitReader 按位流式读取字节切片
// 从高位向低位读
// 注意，可以在每次读取后，判断是否发生错误。也可以在多次读取后，判断是否发生错误。
type BitReader struct {
	core  []byte
	avail uint // 还没有读取的bit数量
	index uint // 从0开始，待读取的字节下标
	pos   uint // 从左往右，从高位往低位 [0, 7]
	err   error
}

func NewBitReader(b []byte) BitReader { _ = "STUB: not implemented"; return *new(BitReader) }

func (br *BitReader) ReadBit() (uint8, error) {
	_ = "STUB: not implemented"
	return 0,

		// @param n: 取值范围 [1, 8]
		nil
}

func (br *BitReader) ReadBits8(n uint) (r uint8, err error) {
	_ = "STUB: not implemented"
	// TODO chef: 8,16,32都去调用ReadBits64会带来额外开销，所以采用实现拷贝的方式，等泛型出来后重构
	return 0, nil
}

// never reach here

// @param n: 取值范围 [1, 16]
func (br *BitReader) ReadBits16(n uint) (r uint16, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// @param n: 取值范围 [1, 32]
func (br *BitReader) ReadBits32(n uint) (r uint32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// @param n: 取值范围 [1, 64]
func (br *BitReader) ReadBits64(n uint) (r uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadBytes
// @param n: 读取多少个字节
func (br *BitReader) ReadBytes(n uint) (r []byte, err error) {
	_ = "STUB: not implemented"
	// 对常见的pos为0的情况单独做优化
	return nil, nil
}

func (br *BitReader) ReadString(n uint) (r string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (br *BitReader) ReadGolomb() (v uint32, err error) {
	_ = "STUB: not implemented"
	return 0,

		// ReadUeGolomb 0阶指数哥伦布编码，无符号
		nil
}

func (br *BitReader) ReadUeGolomb() (v uint32, err error) { _ = "STUB: not implemented"; return 0, nil }

// ReadSeGolomb 哥伦布编码，有符号
func (br *BitReader) ReadSeGolomb() (v int32, err error) { _ = "STUB: not implemented"; return 0, nil }

func (br *BitReader) ReadBits32IgnErr(n uint) uint32 { _ = "STUB: not implemented"; return 0 }

func (br *BitReader) ReadStringIgnErr(n uint) string { _ = "STUB: not implemented"; return "" }

func (br *BitReader) SkipBytes(n uint) error { _ = "STUB: not implemented"; return nil }

func (br *BitReader) SkipBits(n uint) error { _ = "STUB: not implemented"; return nil }

// 返回可读bit数量
func (br *BitReader) AvailBits() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (br *BitReader) Err() error { _ = "STUB: not implemented"; return nil }

func (br *BitReader) readBit() (r uint8, err error) { _ = "STUB: not implemented"; return 0, nil }

// 确保可读空间大小
func (br *BitReader) reserve(n uint) error { _ = "STUB: not implemented"; return nil }

// ----------------------------------------------------------------------------

// TODO chef: BitWriter没有对写越界做检查，由调用方保证这一点，后续可能会加上检查

type BitWriter struct {
	core  []byte
	index int
	pos   uint // 从左往右
}

func NewBitWriter(b []byte) BitWriter { _ = "STUB: not implemented"; return *new(BitWriter) }

// @param b: 当b不为0和1时，取b的最低位
func (bw *BitWriter) WriteBit(b uint8) { _ = "STUB: not implemented"; return }

// 将<v>的低<n>位写入
// @param n: 取值范围 [1, 8]
func (bw *BitWriter) WriteBits8(n uint, v uint8) { _ = "STUB: not implemented"; return }

func (bw *BitWriter) WriteBits16(n uint, v uint16) { _ = "STUB: not implemented"; return }

// ----------------------------------------------------------------------------

// TODO chef: func GetBitX和func GetBitsX没有对写越界做检查，由调用方保证这一点，后续可能会加上检查

// @param pos: 取值范围 [0, 7]，0表示最低位
// @return: [0, 1]
func GetBit8(v uint8, pos uint) uint8 { _ = "STUB: not implemented"; return 0 }

// @param pos: 取值范围 [0, 7]，0表示最低位
// @param n:   取多少位， 取值范围 [1, 8]
//
// 举例，GetBits8(105, 2, 4) = 10（即1010）
//
//	v: 0110 1001
//
// pos:       2
//
//	n:   .. ..
func GetBits8(v uint8, pos uint, n uint) uint8 { _ = "STUB: not implemented"; return 0 }

func GetBit16(v []byte, pos uint) uint8 { _ = "STUB: not implemented"; return 0 }

func GetBits16(v []byte, pos uint, n uint) uint16 { _ = "STUB: not implemented"; return 0 }

var (
	m1 []uint8
)

func init() {
	m1 = []uint8{0, 1, 3, 7, 15, 31, 63, 127, 255} // 0 is dummy
}
