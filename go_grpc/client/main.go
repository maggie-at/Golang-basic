package main

import (
	pb "Golang/go_grpc/server/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接到server端, 此处禁用安全传输, 没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建一个客户端
	client := pb.NewSayHelloClient(conn)

	// client可以直接调用服务端的方法实现
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "client"})
	fmt.Println(resp.GetResponseMsg())
}
