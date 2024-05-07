package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:7890")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte("0D 24 1D 00 80 15 07 16 15 2A 2A 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 8D 0D 0A")); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Message sent to server")

	// 读取服务器回复的消息
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading reply:", err)
		return
	}

	fmt.Println("Reply from server:", string(buffer[:n]))
}
