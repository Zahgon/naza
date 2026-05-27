// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"time"

	"github.com/q191201771/naza/pkg/nazalog"
)

var (
	taskNum       = 1000 * 1000
	initWorkerNum = 1 //1000 * 20 //1000 * 10
)

func originGo() { _ = "STUB: not implemented"; return }

func taskPool() { _ = "STUB: not implemented"; return }

//b.StartTimer()

//b.StopTimer()
//idle, busy := p.Status()
//nazalog.Debugf("done, worker num. idle=%d, busy=%d", idle, busy) // 此时还有个别busy也是正常的，因为只是业务方的任务代码执行完了，可能还没回收到idle队列中
//p.KillIdleWorkers()
//idle, busy = p.Status()
//nazalog.Debugf("killed, worker num. idle=%d, busy=%d", idle, busy)

func main() {
	taskPool()
	//originGo()
	nazalog.Debug("waiting exit.")
	time.Sleep(1000 * time.Second)
	//nazalog.Debug("bye.")
}
