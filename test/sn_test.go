package test

import (
	"testing"
)

func Test_SN(t *testing.T) {
	frame := NewFrame("0D241D0080150716152A2A0100000000000000000000000000008D0D0A")

	// 帧长度 2个字节 即 4个字符
	if frame.GetSN() == "240507990001" {
		t.Log("帧长度校验通过")
	}
}
