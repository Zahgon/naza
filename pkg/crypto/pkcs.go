// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package crypto

import (
	"errors"
)

var ErrPkcs = errors.New("naza.crypto: fxxk")

// @param blockSize 取值范围[0, 255]
//
//	如果是AES，见标准库中aes.BlockSize等于16
func EncryptPkcs7(in []byte, blockSize int) []byte { _ = "STUB: not implemented"; return nil }

func DecryptPkcs7(in []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func EncryptPkcs5(in []byte) []byte { _ = "STUB: not implemented"; return nil }

func DecryptPkcs5(in []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
