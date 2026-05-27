// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"github.com/q191201771/naza/pkg/nazalog"
)

// 分析c++程序pstack的两次结果的差异

type PstackInfo struct {
	tis []ThreadInfo
	tim map[string]ThreadInfo
}

type ThreadInfo struct {
	Num int
	P   string
	Id  int

	RawLine       string
	RawStackLines string
}

func NewPstackInfo(filename string) PstackInfo { _ = "STUB: not implemented"; return *new(PstackInfo) }

//nazalog.Debugf("len(lines)=%d", len(lines))

//nazalog.Debugf("%s", line)

func (pi *PstackInfo) Find(uk string) (ThreadInfo, bool) {
	_ = "STUB: not implemented"
	return *new(ThreadInfo), false
}

func (ti *ThreadInfo) Uk() string { _ = "STUB: not implemented"; return "" }

func parseThreadLine(line string) (num int, p string, id int, err error) {
	_ = "STUB: not implemented"
	return 0, "", 0, nil
}

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.LevelFlag = false
		option.ShortFileFlag = false
		option.TimestampFlag = false
	})

	pi1 := NewPstackInfo("old.txt")
	pi2 := NewPstackInfo("new.txt")
	for _, ti2 := range pi2.tis {
		var pre, suf string
		suf = "\033[0m"
		ti1, exist := pi1.Find(ti2.Uk())
		if exist {
			if ti2.RawStackLines == ti1.RawStackLines {
				// 1, 2都有，但是堆栈没变 没色
				pre = ""
				suf = ""
				nazalog.Debugf("%s-------------------------------------------------------------------------%s", pre, suf)
				nazalog.Debugf("%s%s%s", pre, ti2.RawLine, suf)
				nazalog.Debugf("%s%s%s", pre, ti2.RawStackLines, suf)
			} else {
				// 1, 2都有，但是堆栈变化 红色
				// 注意，断站有变化也可能是函数参数变化了
				pre = "\033[22;31m"
				nazalog.Debugf("%s-------------------------------------------------------------------------%s", pre, suf)
				nazalog.Debugf("%s%s%s", pre, ti2.RawLine, suf)
				nazalog.Debugf("%s%s%s", pre, ti2.RawStackLines, suf)
			}
		} else {
			// 只在2有 绿色
			pre = "\033[22;36m"
			nazalog.Debugf("%s-------------------------------------------------------------------------%s", pre, suf)
			nazalog.Debugf("%s%s%s", pre, ti2.RawLine, suf)
			nazalog.Debugf("%s%s%s", pre, ti2.RawStackLines, suf)
		}
	}
}
