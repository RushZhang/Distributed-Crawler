package rpcsupport

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
	"fmt"
)

//把注册的端口和服务都当做参数传进去(传进去的服务可能是结构体，带有方法的)
func NewServer(host string, service interface{}) error {
	rpc.Register(service)
	fmt.Println("服务器端已经在", host, "端口注册了")
	listener, err := net.Listen("tcp", host)
	if err != nil {panic(err)}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept出现问题： %v", err)
			continue
		}
		// 处理这个链接当然是要go出来一个携程去操作，不耽误无限循环
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}