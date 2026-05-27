// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package filebatch

import (
	"errors"
	"os"
)

// @param path 带路径的文件名
// @param info 文件的 os.FileInfo 信息
// @param content 文件内容
// @return 返回nil或者content原始内容，则不修改文件内容，返回其他内容，则会覆盖重写文件
type WalkFunc func(path string, info os.FileInfo, content []byte, err error) []byte

// 遍历访问指定文件夹下的文件
// @param root 需要遍历访问的文件夹
// @param recursive 是否递归访问子文件夹
// @param suffix 指定文件名后缀进行过滤，如果为""，则不过滤
func Walk(root string, recursive bool, suffix string, walkFn WalkFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// 文件尾部添加内容
func AddTailContent(content []byte, tail []byte) []byte { _ = "STUB: not implemented"; return nil }

// 文件头部添加内容
func AddHeadContent(content []byte, head []byte) []byte { _ = "STUB: not implemented"; return nil }

// 行号范围
// 1表示首行，-1表示最后一行
type LineRange struct {
	From int
	To   int
}

var ErrLineRange = errors.New("naza.filebatch: line range error")

func calcLineRange(len int, lr LineRange) (LineRange, error) {
	_ = "STUB: not implemented"
	// 换算成从0开始的下标
	return *new(LineRange), nil
}

// 排序交换

func DeleteLines(content []byte, lr LineRange) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
