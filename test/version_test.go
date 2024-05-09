package test

import (
	"testing"
)

func Test_Version(t *testing.T) {
	frame := NewFrame("0D241D0080150716152A2A0100000000000000000000000000008D0D0A")

	// 协议版本号 1--255
	if frame.GetVersion() == 1 {
		t.Log("协议版本号校验通过")
	}
}
