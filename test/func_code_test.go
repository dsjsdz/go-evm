package test

import (
	"testing"
)

func Test_Func_Code(t *testing.T) {
	frame := NewFrame("0D241D0080150716152A2A0100000000000000000000000000008D0D0A")

	funcCode := frame.GetFuncCode()
	if funcCode == "80" {
		t.Log("功能码校验通过")
	}
}
