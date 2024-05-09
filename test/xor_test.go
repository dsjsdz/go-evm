package test

import (
	"testing"
)

func Test_Xor(t *testing.T) {
	frame := NewFrame("0D241D0080150716152A2A0100000000000000000000000000008D0D0A")

	// 校验码: 帧长度、功能码、Data异或(即：帧长度^功能码^Data)
	if frame.GetCheckSUM() == "8D" {
		t.Log("校验码校验通过")
	}
}
