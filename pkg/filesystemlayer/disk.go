// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package filesystemlayer

type FslDisk struct {
}

func (f *FslDisk) Type() FslType { _ = "STUB: not implemented"; return *new(FslType) }

func (f *FslDisk) Create(name string) (IFile, error) {
	_ = "STUB: not implemented"
	return *new(IFile), nil
}

func (f *FslDisk) Rename(oldpath string, newpath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FslDisk) MkdirAll(path string, perm uint32) error { _ = "STUB: not implemented"; return nil }

func (f *FslDisk) Remove(name string) error { _ = "STUB: not implemented"; return nil }

func (f *FslDisk) RemoveAll(path string) error { _ = "STUB: not implemented"; return nil }

func (f *FslDisk) ReadFile(filename string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FslDisk) WriteFile(filename string, data []byte, perm uint32) error {
	_ = "STUB: not implemented"
	return nil
}
