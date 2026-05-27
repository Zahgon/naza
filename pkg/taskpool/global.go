// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package taskpool

var global Pool

func Go(task TaskFn, param ...interface{}) { _ = "STUB: not implemented"; return }

func GetCurrentStatus() Status { _ = "STUB: not implemented"; return *new(Status) }

func KillIdleWorkers() { _ = "STUB: not implemented"; return }

func Init(modOptions ...ModOption) error { _ = "STUB: not implemented"; return nil }

func init() {
	_ = Init()
}
