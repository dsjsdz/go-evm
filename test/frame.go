package test

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

type Frame struct {
	Body string
}

func NewFrame(body string) *Frame {
	return &Frame{
		Body: body,
	}
}

// GetSN 设备ID 14个字节
func (f *Frame) GetSN() string {
	var (
		payload = f.Body[10:38] // 10: 10+14*2
		sn      string
		l       = len(payload) / 2
	)

	// 用户自定义判断
	if f.GetFuncCode() != "80" {
		return "非法功能码"
	}

	// 转成10进制内容
	for i := 0; i < (l - 1); i++ {
		var (
			start = i * 2
			end   = i*2 + 2
		)

		temp := payload[start:end]
		value, err := strconv.ParseInt(temp, 16, 64)
		if err != nil {
			continue
		}

		// ASCII码
		num := value - '0'
		if num < 0 {
			// 存在 00 为 -48 不是合法内容
			break
		}

		sn += fmt.Sprintf("%v", num)
	}

	return ""
}

// GetVersion 协议版本号 1--255
func (f *Frame) GetVersion() int64 {
	version, err := strconv.ParseInt(f.Body[38:40], 16, 64)
	if err != nil {
		return 0
	}

	return version
}

func (f *Frame) GetHeader() string {
	return f.Body[0:4]
}

func (f *Frame) GetFooter() string {
	var (
		l = len(f.Body)
	)
	return f.Body[l-4 : l]
}

func (f *Frame) GetData() string {
	var (
		l = len(f.Body)
	)
	return f.Body[10 : l-6]
}

// GetLength 低位在前，高位在后
func (f *Frame) GetLength() int64 {
	size, _ := strconv.ParseInt(f.Body[4:6], 16, 64)
	return size
}

// GetFuncCode 1个字节 内容(十六进制) 即: 2个字符
func (f *Frame) GetFuncCode() string {
	return f.Body[8:10]
}

func (f *Frame) GetCheckSUM() string {
	var (
		payload = []string{
			f.Body[4:8],                // 帧长度 2个字节
			f.Body[8:10],               // 功能码 1个字节
			f.Body[10 : len(f.Body)-6], // Data N个字节
		}
		bbc byte
	)

	for _, v := range payload {
		bytes, _ := hex.DecodeString(v)
		for _, b := range bytes {
			bbc ^= b
		}
	}

	return fmt.Sprintf("%02X", bbc)
}
