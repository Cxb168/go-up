package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		panic(err)
	}
	go func() {
		sendIndex := 0
		for {
			sendIndex ++
			sendMsg := fmt.Sprintf("Msg: [%d]...... \n", sendIndex)
			//sendMsg := fmt.Sprintf("Msg %d", sendIndex)
			_, sendErr := conn.Write([]byte(sendMsg))
			if sendErr != nil {
				panic(sendErr)
			}
			fmt.Println(sendMsg)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			receive, rErr := io.ReadAll(conn)
			if rErr != nil {
			    panic(rErr)
			}
			fmt.Printf("Client revice %s\n", receive)
		}
	}()

	select {
	}
}
