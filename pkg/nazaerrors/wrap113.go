// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

//go:build go1.13
// +build go1.13

package nazaerrors

func Wrap(err error, msg ...string) error { _ = "STUB: not implemented"; return nil }

// TODO(chef): 整理下面三个函数

func Unwrap(err error) error { _ = "STUB: not implemented"; return nil }

func Is(err, target error) bool { _ = "STUB: not implemented"; return false }

func As(err error, target interface{}) bool { _ = "STUB: not implemented"; return false }
