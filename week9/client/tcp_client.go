package main

import (
	"fmt"
	protocol2 "go-up/week9/protocol"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		panic(err)
	}
	proto := protocol2.NewProtocol(conn)
	go func() {
		sendIndex := 0
		for {
			sendIndex++
			//sendMsg := fmt.Sprintf("%d......../........./........./..", sendIndex)
			sendMsg := fmt.Sprintf("%d...", sendIndex)
			_, sendErr := conn.Write(proto.Encode([]byte(sendMsg)))
			if sendErr != nil {
				panic(sendErr)
			}
			//fmt.Println(sendMsg)
			time.Sleep(time.Second)
		}
	}()

	go proto.Run()
	for info := range proto.GetChan() {
		fmt.Printf("Client receive: %s\n", info.Body)
	}
}
