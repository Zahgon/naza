// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

// Package bininfo
//
// 将编译时源码的git版本信息（当前tag，commit log的sha值和commit message，是否有未提交的修改），编译时间，Go版本，编译、运行平台打入程序中
// 编译时传入这些信息的方式见 naza 的编译脚本： https://github.com/q191201771/naza/blob/master/build.sh
package bininfo

var (
	// 初始化为 unknown，如果编译时没有传入这些值，则为 unknown
	GitTag         = "unknown"
	GitCommitLog   = "unknown"
	GitStatus      = "unknown"
	BuildTime      = "unknown"
	BuildGoVersion = "unknown"
)

// 返回单行格式
func StringifySingleLine() string { _ = "STUB: not implemented"; return "" }

// 返回多行格式
func StringifyMultiLine() string { _ = "STUB: not implemented"; return "" }

// 对一些值做美化处理
func beauty() { _ = "STUB: not implemented"; return }

// GitStatus 为空时，说明本地源码与最近的 commit 记录一致，无修改
// 为它赋一个特殊值

// 将多行结果合并为一行

func init() {
	beauty()
}
