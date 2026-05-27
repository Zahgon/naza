// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazareflect

func IsNil(actual interface{}) bool { _ = "STUB: not implemented"; return false }

// TODO chef: 考虑是否将EqualInteger放入Equal中，但是需考虑，会给Equal带来额外的性能开销
func Equal(expected, actual interface{}) bool { _ = "STUB: not implemented"; return false }

// 用于判断不同类型的整型的值是否相等，比如int8(1)和int32(1)做比较，结果为true
// 注意，如果a或b不是整型，则直接返回false
func EqualInteger(a, b interface{}) bool {
	_ = "STUB: not implemented"
	// a有3种状态：有符号整型，无符号整型，非整型
	// b同理也是3种状态，
	// 那么总共有3*3种组合需要判断
	return false
}

// a,b都是有符号整型 (1)

// a,b都是无符号整型 (1)

// a是有符号整型，b是无符号整型 (1)

// a是无符号整型，b是有符号整型 (1)

// 剩下的情况，至少有一个不是整型 (5)

func tryInt(actual interface{}) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func tryUint(actual interface{}) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }
