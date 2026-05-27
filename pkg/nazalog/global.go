// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazalog

var global Logger

func Tracef(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Panicf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Trace(v ...interface{}) { _ = "STUB: not implemented"; return }

func Debug(v ...interface{}) { _ = "STUB: not implemented"; return }

func Info(v ...interface{}) { _ = "STUB: not implemented"; return }

func Warn(v ...interface{}) { _ = "STUB: not implemented"; return }

func Error(v ...interface{}) { _ = "STUB: not implemented"; return }

func Fatal(v ...interface{}) { _ = "STUB: not implemented"; return }

func Panic(v ...interface{}) { _ = "STUB: not implemented"; return }

func Output(calldepth int, s string) error { _ = "STUB: not implemented"; return nil }

func Print(v ...interface{}) { _ = "STUB: not implemented"; return }

func Printf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Println(v ...interface{}) { _ = "STUB: not implemented"; return }

func Fatalln(v ...interface{}) { _ = "STUB: not implemented"; return }

func Panicln(v ...interface{}) { _ = "STUB: not implemented"; return }

func Assert(expected interface{}, actual interface{}, extInfo ...string) {
	_ = "STUB: not implemented"
	return
}

func Out(level Level, calldepth int, s string) { _ = "STUB: not implemented"; return }

func Sync() { _ = "STUB: not implemented"; return }

func WithPrefix(s string) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func GetOption() Option {
	_ = "STUB: not implemented"
	return *

	// ---------------------------------------------------------------------------------------------------------------------
	new(Option)
}

// GetGlobalLogger 获取全局Logger
func GetGlobalLogger() Logger {
	_ = "STUB: not implemented"

	// Init 初始化全局Logger
	//
	// 注意，全局Logger在不需要特殊配置时，可以不显示调用 Init 函数
	// 注意，该方法不会修改global指针指向，而是操作global指针指向的对象
	return *new(Logger)
}

func Init(modOptions ...ModOption) error { _ = "STUB: not implemented"; return nil }

// SetGlobalLogger 更换全局Logger
//
// 注意，更换后，之前调用 GetGlobalLogger 获取的全局Logger和当前的全局Logger将是两个对象
//
// TODO(chef): [refactor] 在已经提供 Init 的前提下，是否应该删除掉该函数
func SetGlobalLogger(l Logger) {
	_ = "STUB: not implemented"

	// ---------------------------------------------------------------------------------------------------------------------
	return
}

func init() {
	global, _ = newLogger()
}
