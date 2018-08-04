package main

import (
	"net/rpc"
	"crawler/rpcTest"
	"net"
	"log"
	"net/rpc/jsonrpc"
	"fmt"
)

func main() {
	rpc.Register(rpcTest.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {panic(err)}
	fmt.Println("服务器已经在1234端口监听。。。")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept出现问题： %v", err)
			continue
		}
		// 处理这个链接当然是要go出来一个携程去操作，不耽误无限循环
		go jsonrpc.ServeConn(conn)
	}
}
