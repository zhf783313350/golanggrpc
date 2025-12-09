package main

import (
	"context"
	"fmt"
	pb "golanggrpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 故意写错字段名,引入编译错误
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	// 创建 gRPC 服务器
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterHelloServiceServer(s, &Server{}) // ✅ 正确!
	fmt.Println("gRPC 服务器启动,监听端口 :50051")
	// 启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
