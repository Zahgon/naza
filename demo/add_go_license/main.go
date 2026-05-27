// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/q191201771/naza/pkg/filebatch"
	"github.com/q191201771/naza/pkg/nazalog"
)

var licenseTmpl = `// Copyright %d, %s.  All rights reserved.
// https://%s
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: %s (%s)

`

func main() {
	dir, name, email := parseFlag()

	year := time.Now().Year()
	repo := achieveRepo(dir)
	license := fmt.Sprintf(licenseTmpl, year, name, repo, name, email)
	nazalog.Debug(license)

	var (
		skipCount int
		modCount  int
	)
	err2 := filebatch.Walk(dir, true, ".go", func(path string, info os.FileInfo, content []byte, err error) []byte {
		if err != nil {
			nazalog.Warnf("read file failed. file=%s, err=%+v", path, err)
			return nil
		}
		lines := bytes.Split(content, []byte{'\n'})
		if bytes.Index(lines[0], []byte("Copyright")) != -1 {
			skipCount++
			//nc, _ := filebatch.DeleteLines(content, filebatch.LineRange{From:1, To:7})
			//return nc
			return nil
		}

		modCount++
		return filebatch.AddHeadContent(content, []byte(license))
	})
	nazalog.Assert(nil, err2)
	nazalog.Infof("count. mod=%d, skip=%d", modCount, skipCount)
}

func achieveRepo(root string) string { _ = "STUB: not implemented"; return "" }

func parseFlag() (string, string, string) { _ = "STUB: not implemented"; return "", "", "" }
