// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package filesystemlayer

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("naza filesystemlayer: not found")

type FslMemory struct {
	mu    sync.Mutex
	files map[string]*file // key filename
}

type file struct {
	buf []byte
}

func NewFslMemory() *FslMemory { _ = "STUB: not implemented"; return nil }

func (f *FslMemory) Type() FslType { _ = "STUB: not implemented"; return *new(FslType) }

func (f *FslMemory) Create(name string) (IFile, error) {
	_ = "STUB: not implemented"
	return *new(IFile), nil
}

func (f *FslMemory) Rename(oldpath string, newpath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FslMemory) MkdirAll(path string, perm uint32) error { _ = "STUB: not implemented"; return nil }

func (f *FslMemory) Remove(name string) error { _ = "STUB: not implemented"; return nil }

func (f *FslMemory) RemoveAll(path string) error { _ = "STUB: not implemented"; return nil }

func (f *FslMemory) ReadFile(filename string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FslMemory) WriteFile(filename string, data []byte, perm uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (f *FslMemory) openFile(name string, flag int, perm uint32) (IFile, error) {
	_ = "STUB: not implemented"
	return *new(IFile), nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (f *file) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (f *file) Close() error { _ = "STUB: not implemented"; return nil }

func (f *file) truncate() { _ = "STUB: not implemented"; return }

func (f *file) clone() []byte { _ = "STUB: not implemented"; return nil }
