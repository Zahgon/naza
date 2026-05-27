// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazahttp

// TODO(chef): 重命名为GetAll
//
// GetHttpFile 获取http文件保存至字节切片
func GetHttpFile(url string, timeoutMs int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 获取http文件保存至本地
func DownloadHttpFile(url string, saveTo string, timeoutMs int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
