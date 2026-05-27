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

// @param url  地址
// @param info 需要序列化的结构体
// @param client 注意，如果为nil，则使用http.DefaultClient
func PostJson(url string, info interface{}, client *http.Client) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
