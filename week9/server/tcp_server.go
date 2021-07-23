package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("Socket server")
	listener, err := net.Listen("tcp", "0.0.0.0:8899")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		fmt.Println("------Accept-----", conn.RemoteAddr())
		if err != nil {
			panic(err)
		}
		go func() {
			for{
				receiveBytes := make([]byte, 64)
				_, readErr := io.ReadFull(conn, receiveBytes)
				if readErr != nil {
					panic(readErr)
				}
				fmt.Println("------------", string(receiveBytes))
				r := fmt.Sprintf("Recove by[%s]", receiveBytes)
				_, wErr := conn.Write([]byte(r))
				if wErr != nil {
				    panic(wErr)
				}
			}
		}()
	}
}
