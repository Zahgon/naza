// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazahttp

import (
	"io"
	"net/http"
)

// e.g. bufio.Reader
type HttpReader interface {
	LineReader
	io.Reader
}

type HttpMsgCtx struct {
	ReqMethodOrRespVersion string
	ReqUriOrRespStatusCode string
	ReqVersionOrRespReason string
	Headers                http.Header
	Body                   []byte
}

type HttpReqMsgCtx struct {
	Method  string
	Uri     string
	Version string
	Headers http.Header
	Body    []byte
}

type HttpRespMsgCtx struct {
	Version    string
	StatusCode string
	Reason     string
	Headers    http.Header
	Body       []byte
}

func ReadHttpRequestMessage(r HttpReader) (ctx HttpReqMsgCtx, err error) {
	_ = "STUB: not implemented"
	return *new(HttpReqMsgCtx), nil
}

func ReadHttpResponseMessage(r HttpReader) (ctx HttpRespMsgCtx, err error) {
	_ = "STUB: not implemented"
	return *new(HttpRespMsgCtx), nil
}

// ReadHttpMessage
//
// 注意，如果HTTP Header中不包含`Content-Length`，则不会读取HTTP Body，并且err返回值为nil
func ReadHttpMessage(r HttpReader) (ctx HttpMsgCtx, err error) {
	_ = "STUB: not implemented"
	return *new(HttpMsgCtx), nil
}
