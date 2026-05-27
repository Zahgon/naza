// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"fmt"
	"time"
)

type Origin struct {
	a uint64
	b uint64
}

type WithPadding struct {
	a uint64
	_ [56]byte
	b uint64
	_ [56]byte
}

var num = 1000 * 1000

func OriginParallel() { _ = "STUB: not implemented"; return }

func WithPaddingParallel() { _ = "STUB: not implemented"; return }

func main() {
	var b time.Time

	b = time.Now()
	OriginParallel()
	fmt.Printf("OriginParallel. Cost=%+v.\n", time.Now().Sub(b))

	b = time.Now()
	WithPaddingParallel()
	fmt.Printf("WithPaddingParallel. Cost=%+v.\n", time.Now().Sub(b))
}
