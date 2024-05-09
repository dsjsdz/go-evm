package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"log"
	"strings"
)

type echoServer struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool
}

func (es *echoServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("echo server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
	return gnet.None
}

func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
	var (
		buf, _ = c.Next(-1)
	)

	raw := hex.EncodeToString(buf)  // 16进字 接收
	payload := strings.ToUpper(raw) // 转成大写

	fmt.Printf("接收到内容: %v", payload)
	// 0D2424008032343035303839393030303100000100000000000000000000000000AF0D0A

	// TODO: 根据功能码进行判断，并处理对应逻辑

	// 回传内容 自定义 根据文档
	go c.Write([]byte("0D2424008032343035303839393030303100000100000000000000000000000000AF0D0A"))
	return gnet.None
}

func main() {
	var port int
	var multicore bool

	// Example command: go run echo.go --port 7890 --multicore=true
	flag.IntVar(&port, "port", 7891, "--port 7891")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()
	echo := &echoServer{addr: fmt.Sprintf("tcp://:%d", port), multicore: multicore}
	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(multicore)))
}
