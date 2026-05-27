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

// @brief 从http请求中解析body中的json字符串，并反序列化至结构体中
//
// @param r            http请求对象
// @param info         输出参数，用于接收反序列化之后的数据
// @param keyFieldList 可选参数，可指定一个或多个json中必须存在的字段
func UnmarshalRequestJsonBody(r *http.Request, info interface{}, keyFieldList ...string) error {
	_ = "STUB: not implemented"
	return nil
}
