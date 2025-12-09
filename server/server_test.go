package main

import (
    "context"
    "testing"
    pb "golanggrpc/proto"
)

func TestSayHello(t *testing.T) {
    // 创建服务器实例
    s := &Server{}
    
    // 测试数据
    req := &pb.HelloRequest{Name: "测试"}
    
    // 调用方法
    resp, err := s.SayHello(context.Background(), req)
    
    // 验证结果
    if err != nil {
        t.Fatalf("SayHello 失败: %v", err)
    }
    
    expected := "Hello 测试"
    if resp.Message != expected {
        t.Errorf("期望 %s, 得到 %s", expected, resp.Message)
    }
}