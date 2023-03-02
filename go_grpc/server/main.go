package main

import (
	pb "Golang/go_grpc/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// server 实现了SayHelloServer接口
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println(req.RequestName + " is calling")
	return &pb.HelloResponse{ResponseMsg: "hello~ " + req.RequestName}, nil
}

func main() {
	// 开启端口监听
	listen, _ := net.Listen("tcp", ":8888") // 注意冒号
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 在grpc服务端去注册server这个服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	err := grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
