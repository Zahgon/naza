// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazasync

import (
	"errors"
)

// NOTICE copy from https://github.com/golang/net/blob/master/http2/gotrack.go

var ErrObtainGoroutineId = errors.New("nazasync: obtain current goroutine id failed")

func CurGoroutineId() (int64, error) { _ = "STUB: not implemented"; return 0, nil }
