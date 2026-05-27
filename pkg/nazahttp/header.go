// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazahttp

import (
	"net/http"
)

type LineReader interface {
	ReadLine() (line []byte, isPrefix bool, err error)
}

// ReadHttpHeader
//
// @return firstLine: request的request line或response的status line
// @return headers:   request header fields的键值对
func ReadHttpHeader(r LineReader) (firstLine string, headers http.Header, err error) {
	_ = "STUB: not implemented"
	return "", *new(http.Header), nil
}

// 读到一个空的 \r\n 表示http头全部读取完毕了

// 兼容性处理，见单元测试TestReadHttpResponseMessage中的case1
//
// 如果找不到冒号，就把它算到上一个header条目的value里
// 也即我们认为上一个value中的自身内容包含了格式错误的\r\n
//

// Request-Line = Method SP URI SP Version CRLF
func ParseHttpRequestLine(line string) (method string, uri string, version string, err error) {
	_ = "STUB: not implemented"
	return "", "",

		// Status-Line = Version SP Status-Code SP Reason CRLF
		"", nil
}

func ParseHttpStatusLine(line string) (version string, statusCode string, reason string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

func parseFirstLine(line string) (item1, item2, item3 string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

// TODO(chef): refactor 整理此处代码，可使用split根据数量返回

//if s == -1 || f+1+s+1 == len(line) {
//	err = nazaerrors.Wrap(ErrFirstLine, line)
//	return
//}
