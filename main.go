package main

import (
	"flag"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"log"
	"time"
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
		now    = time.Now().Format(time.DateTime)
	)
	fmt.Printf("[%s] received: %v \n", now, string(buf))

	//payload := strings.Split(string(buf), " ")
	//if len(payload) == 0 {
	//	log.Println("桢内容格式不正确")
	//	return gnet.None
	//}
	//
	//// 帧长度 即 index 为 2-3 的内容，3默认是 00，不去计算
	//size, err := strconv.ParseInt(payload[util.FrameLenSize], 16, 64)
	//if err != nil {
	//	log.Println(err)
	//	return gnet.None
	//}
	//
	//if len(payload) != int(size) {
	//	log.Println("桢内容长度不正确")
	//	return gnet.None
	//}
	//
	//var (
	//	funcCode = payload[4]
	//	reply    []byte
	//	now      = time.Now()
	//	year     = strconv.FormatInt(int64(now.Year()), 16)
	//	month    = strconv.FormatInt(int64(now.Month()), 16)
	//	day      = strconv.FormatInt(int64(now.Day()), 16)
	//	hour     = strconv.FormatInt(int64(now.Hour()), 16)
	//	minute   = strconv.FormatInt(int64(now.Minute()), 16)
	//	second   = strconv.FormatInt(int64(now.Second()), 16)
	//	tmpl     = "0D 24 1D 00 80 %v %v %v %v %v %v %s 00 00 00 00 00 00 00 00 00 00 00 00 00 00 8D 0D 0A"
	//)
	//switch funcCode {
	//case util.MethodServerTimeGet: // 登陆服务器
	//	// NOTE: 处理逻辑，FrameMachineIdSize 设备id长度14
	//	reply = []byte(fmt.Sprintf(tmpl, year, month, day, hour, minute, second, util.SuccessStatusCode))
	//default:
	//	// 其他自行处理
	//}
	//
	//go c.Write(reply)
	return gnet.None
}

func main() {
	var port int
	var multicore bool

	// Example command: go run echo.go --port 7890 --multicore=true
	flag.IntVar(&port, "port", 7890, "--port 7890")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()
	echo := &echoServer{addr: fmt.Sprintf("tcp://:%d", port), multicore: multicore}
	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(multicore)))
}
