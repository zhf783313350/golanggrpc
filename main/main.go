package main

import (
	"context"
	"fmt"
	pb "golanggrpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
    // 1. 连接服务器
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("连接失败: %v", err)
    }
    defer conn.Close()
    
    // 2. 创建客户端
    c := pb.NewHelloServiceClient(conn)
    
    // 3. 设置超时
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    
    // 4. 调用方法
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "世界"})
    
    // 5. 打印结果
    fmt.Printf("响应: %s\n", r.GetMessage())
}