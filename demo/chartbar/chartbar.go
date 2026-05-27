// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"fmt"

	"github.com/q191201771/naza/pkg/chartbar"

	"github.com/q191201771/naza/pkg/nazalog"
)

func main() {
	filename := parseFlag()
	output, err := chartbar.DefaultCtx.WithCsv(filename)
	nazalog.Assert(nil, err)
	fmt.Print(output)
}

func parseFlag() string { _ = "STUB: not implemented"; return "" }
