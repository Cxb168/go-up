package main

import (
	"fmt"
	"net/http"
	"time"
)

//1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

type handle struct {
	srv Server
}

func (h *handle) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(fmt.Sprintf("Server %s health, %s\n", h.srv.Name(), time.Now().String())))
}

type HttpServer struct {
	name string
	addr string
}

func NewHttpServer(name, addr string) Server {
	return HttpServer{
		name: name,
		addr: addr,
	}
}

func (s HttpServer) Start() error {
	fmt.Printf("server %s start at %s\n", s.name, s.addr)
	handle := &handle{srv: s}
	return http.ListenAndServe(s.addr, handle)
}

func (s HttpServer) Stop() error {
	return nil
}

func (s HttpServer) Name() string {
	return s.name
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
