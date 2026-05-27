// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"bytes"
	"os"
	"strings"

	"github.com/q191201771/naza/pkg/filebatch"
	"github.com/q191201771/naza/pkg/nazalog"
)

// 帮助找出源码中多个大写字母连接在一起的地方

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.Level = nazalog.LevelInfo
		//option.LevelFlag = false
		option.TimestampWithMsFlag = false
		option.TimestampFlag = false
		//option.ShortFileFlag = false
	})

	dir := parseFlag()

	// 遍历所有go文件
	err := filebatch.Walk(dir, true, ".go", func(path string, info os.FileInfo, content []byte, err error) []byte {
		if err != nil {
			nazalog.Warnf("read file failed. file=%s, err=%+v", path, err)
			return nil
		}

		nazalog.Tracef("path:%s", path)

		//checkModFile(path, content)
		//return nil

		//免检的文件：
		if strings.Contains(path, "/pkg/alpha/stun/") {
			return nil
		}

		// 免检文件：测试文件
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		lines := bytes.Split(content, []byte{'\n'})

		// 免检的行：
		ignContainsKeyList := []string{
			// 字符串
			"\"",
			// 16进制的数字
			"0x",
			// 接口
			"IBufWriter",
			"IClientSession",
			"IServerSession",
			"IClientSessionLifecycle",
			"IServerSessionLifecycle",
			"ISessionStat",
			"ISessionUrlContext",
			"IObject",
			"IPathStrategy",
			"IPathRequestStrategy",
			"IPathWriteStrategy",
			"IQueueObserver",
			"IHandshakeClient",
			"IRtpUnpacker",
			"IRtpUnpackContainer",
			"IRtpUnpackerProtocol",
			"IInterleavedPacketWriter",
			"filesystemlayer.IFileSystemLayer",
			"filesystemlayer.IFile",
			"IFile",
			"IFileSystemLayer",
			// RTSP相关
			"HeaderCSeq",
			"ARtpMap",
			"AFmtPBase",
			"AControl",
			// 标准库
			".URL",
			".TLS",
			".SIGUSR",
			"ServeHTTP(",
			".URI",
			".RequestURI",
			"io.EOF",
			"net.UDPAddr",
			"net.UDPConn",
			"net.ResolveUDPAddr",
			"net.ListenUDP",
			"WriteToUDP",
			"ReadFromUDP",
			"time.RFC1123",
			"runtime.GOOS",
			"crc32.ChecksumIEEE",
			"cipher.NewCBCEncrypter",
			"cipher.NewCBCDecrypter",
			"os.O_CREATE",
			//
			"LAddr",
			"RAddr",
		}

		// 注释
		ignPrefixKeyList := []string{
			"//",
			"/*",
		}

		// 逐行分析
		for j, line := range lines {
			// 免检的行：
			if j == 3 && strings.Contains(string(line), "MIT-style license") {
				continue
			}

			ignFlag := false
			for _, k := range ignContainsKeyList {
				if strings.Contains(string(line), k) {
					nazalog.Debugf("ign contains line:%s %s", string(line), k)
					ignFlag = true
				}
			}
			if ignFlag {
				continue
			}

			for _, k := range ignPrefixKeyList {
				if strings.HasPrefix(strings.TrimSpace(string(line)), k) {
					nazalog.Debugf("ign prefix line:%s %s", string(line), k)
					ignFlag = true
				}
			}
			if ignFlag {
				continue
			}

			for i := range line {
				// 连续两个字符是大写字母
				if i == 0 {
					continue
				}
				if isCap(line[i]) && isCap(line[i-1]) {
					nazalog.Infof("%s:%d %s", path, j+1, string(highlightSerialCap(line)))
					break
				}
			}
		}
		return nil
	})
	nazalog.Assert(nil, err)
}

// 是否大写字母
func isCap(c byte) bool { _ = "STUB: not implemented"; return false }

func min(a, b int) int { _ = "STUB: not implemented"; return 0 }

// 有连续大写的地方高亮显示
func highlightSerialCap(line []byte) []byte { _ = "STUB: not implemented"; return nil }

// 一段检查文件修改后和修改前的逻辑，修改是否符合预期
func checkModFile(path string, content []byte) { _ = "STUB: not implemented"; return }

// 理论上，大部分修改，不影响文件大小

// 不管大小是否相等，取最小值，逐个字节比较内容
// 理论上，大部分修改，要么是相等，要么是将新内容从小写转换回大写就相等

func parseFlag() string { _ = "STUB: not implemented"; return "" }
