package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

//1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

type HttpServer struct {
	name   string
	addr   string
	server *http.Server
}

func NewHttpServer(name, addr string) Server {
	return &HttpServer{
		name: name,
		addr: addr,
	}
}

func (s *HttpServer) Start() error {
	server := &http.Server{Addr: s.addr}
	s.server = server
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fmt.Sprintf("Server %s health, %s\n", s.name, time.Now().String())))
	})
	fmt.Printf("server %s start at %s\n", s.name, s.addr)
	server.Handler = serverMux
	return server.ListenAndServe()
}

func (s *HttpServer) Stop() error {
	return s.server.Shutdown(context.Background())
}

func main() {
	s1 := NewHttpServer("Server1", "0.0.0.0:8081")
	s2 := NewHttpServer("Server2", "0.0.0.0:8082")

	app := NewApp(
		WithName("App1"),
		WithServer(s1, s2),
	)
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
