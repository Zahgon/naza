// Copyright 2024, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

//go:build windows
// +build windows

package nazalog

func (l *logger) writeLevelStringIfNeeded(level Level) { _ = "STUB: not implemented"; return }

// windows系统不用写带颜色的日志级别字段
