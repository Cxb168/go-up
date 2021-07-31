package main

import (
	"fmt"
	protocol2 "go-up/week9/protocol"
	"net"
	"time"
)

// 协议解析失败，如何处理（报错并关闭连接）
// TODO 支持 json 与 Protobuffer； 确保送达
func main() {
	fmt.Println("Socket server")
	listener, err := net.Listen("tcp", "0.0.0.0:8899")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		//conn.SetReadDeadline()
		fmt.Println("------Accept-----", conn.RemoteAddr())
		if err != nil {
			panic(err)
		}
		//go func() {
		proto := protocol2.NewProtocol(conn)
		go proto.Run()
		func() {
			for info := range proto.GetChan() {
				fmt.Printf("Receive Msg: %s\n", info.Body)
				// TODO 根据收到消息，调用接口，返回数据
				rtn := fmt.Sprintf("Ack for %s", info.Body)
				_, err = conn.Write(proto.Encode([]byte(rtn)))
				if err != nil {
				    panic(err)
				}
			}
		}()
		//}()
	}
}

func tryWrite2Client(conn net.Conn, proto protocol2.Protocol) {
	index := 0
	for {
		index += 1
		msg := fmt.Sprintf("Server msg %d", index)
		_, err := conn.Write(proto.Encode([]byte(msg)))
		if err != nil {
			fmt.Println("Error:", time.Now())
			panic(err)
		}
		fmt.Println(time.Now(), msg)
		time.Sleep(time.Second * 5)
	}
}
