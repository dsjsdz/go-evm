package test

import (
	"fmt"
	"net"
	"testing"
)

func Test_Client(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:7890")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte("0D241D0080150716152A2A0100000000000000000000000000008D0D0A")); err != nil {
		t.Log(err)
		return
	}
	t.Log("Message sent to server")

	// 读取服务器回复的消息
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		t.Log("Error reading reply:", err)
		return
	}

	t.Log("Reply from server:", string(buffer[:n]))
}
