package main

import (
	"context"
	"fmt"
	"go-away/service"
	"google.golang.org/grpc"
)

func main(){
	// 连接服务器
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	defer conn.Close()

	// 新建一个客户端
	c := service.NewGreeterClient(conn)

	// 调用服务端函数
	r, err := c.SayHello(context.Background(), &service.HelloRequest{Name: "llf"})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}

	fmt.Printf("调用成功: %s", r.Message)
}
