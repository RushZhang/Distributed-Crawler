package main

import (
	"net"
	"net/rpc/jsonrpc"
	"crawler/rpcTest"
	"fmt"
)


//不同于在命令行手动去链接服务器, 这里自动连接
//telnet localhost 1234
//{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64

	err = client.Call("DemoService.Div", rpcTest.Args{10, 3}, &result)
	fmt.Println(result, err)

	err = client.Call("DemoService.Div", rpcTest.Args{10, 0}, &result)
	fmt.Println(result, err)
}
