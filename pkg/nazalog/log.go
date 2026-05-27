// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazalog

import (
	"bytes"
	"os"
	"sync"
	"time"

	"github.com/q191201771/naza/pkg/mock"

	"github.com/q191201771/naza/pkg/nazacolor"
)

var _ Logger = new(logger)

var Clock = mock.NewStdClock()

const (
	levelTraceString = "TRACE "
	levelDebugString = "DEBUG "
	levelInfoString  = " INFO "
	levelWarnString  = " WARN "
	levelErrorString = "ERROR "
	levelFatalString = "FATAL "
	levelPanicString = "PANIC "

	levelTraceColorString = nazacolor.SimplePrefixGreen + levelTraceString + nazacolor.SimpleSuffix
	levelDebugColorString = nazacolor.SimplePrefixBlue + levelDebugString + nazacolor.SimpleSuffix
	levelInfoColorString  = nazacolor.SimplePrefixCyan + levelInfoString + nazacolor.SimpleSuffix
	levelWarnColorString  = nazacolor.SimplePrefixYellow + levelWarnString + nazacolor.SimpleSuffix
	levelErrorColorString = nazacolor.SimplePrefixRed + levelErrorString + nazacolor.SimpleSuffix
	levelFatalColorString = nazacolor.SimplePrefixRed + levelFatalString + nazacolor.SimpleSuffix
	levelPanicColorString = nazacolor.SimplePrefixRed + levelPanicString + nazacolor.SimpleSuffix
)

var (
	levelToString = map[Level]string{
		LevelTrace: levelTraceString,
		LevelDebug: levelDebugString,
		LevelInfo:  levelInfoString,
		LevelWarn:  levelWarnString,
		LevelError: levelErrorString,
		LevelFatal: levelFatalString,
		LevelPanic: levelPanicString,
	}
	levelToColorString = map[Level]string{
		LevelTrace: levelTraceColorString,
		LevelDebug: levelDebugColorString,
		LevelInfo:  levelInfoColorString,
		LevelWarn:  levelWarnColorString,
		LevelError: levelErrorColorString,
		LevelFatal: levelFatalColorString,
		LevelPanic: levelPanicColorString,
	}
)

type logger struct {
	prefixs []string
	core    *core
}

type core struct {
	option Option

	m             sync.Mutex
	fp            *os.File
	console       *os.File
	buf           bytes.Buffer // TODO(chef): [refactor] 是否需要使用nazabytes.Buffer
	currRoundTime time.Time
}

func (l *logger) Tracef(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Panicf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Trace(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Debug(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Info(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Warn(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Error(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Fatal(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Panic(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Output(calldepth int, s string) error { _ = "STUB: not implemented"; return nil }

func (l *logger) Print(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Printf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Println(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Fatalln(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Panicln(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Assert(expected interface{}, actual interface{}, extInfo ...string) {
	_ = "STUB: not implemented"
	return
}

func (l *logger) Out(level Level, calldepth int, s string) { _ = "STUB: not implemented"; return }

// 输出至控制台

// 输出至日志文件

// 同时满足条件时，翻滚一次就够了

// 忽略关闭的错误

// windows会走这个逻辑分支
// TODO(chef): 应判断具体的错误值 202302

// 输出至hook

func (l *logger) Sync() { _ = "STUB: not implemented"; return }

func (l *logger) WithPrefix(s string) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (l *logger) GetOption() Option { _ = "STUB: not implemented"; return *new(Option) }

func (l *logger) Init(modOptions ...ModOption) error { _ = "STUB: not implemented"; return nil }

// ---------------------------------------------------------------------------------------------------------------------

func newLogger(modOptions ...ModOption) (*logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validate(option Option) error { _ = "STUB: not implemented"; return nil }

func writeTime(buf *bytes.Buffer, t time.Time, withMs bool) { _ = "STUB: not implemented"; return }

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// @NOTICE 该函数拷贝自 Go 标准库 /src/log/log.go: func itoa(buf *[]byte, i int, wid int)
// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
func itoa(buf *bytes.Buffer, i int, wid int) {
	_ = "STUB: not implemented"
	// Assemble decimal in reverse order.
	return
}

// i < 10
