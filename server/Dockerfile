# 第一阶段: 构建阶段
FROM golang:alpine AS builder

# 安装必要工具
RUN apk add --no-cache git protobuf protobuf-dev

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

# 安装 protoc 插件
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 复制整个项目
COPY . .

# 生成 proto 代码
RUN export PATH="$PATH:$(go env GOPATH)/bin" && \
    protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/hello.proto

# 构建服务端
WORKDIR /app/server
RUN CGO_ENABLED=0 GOOS=linux go build -o server-app .

# 第二阶段: 运行阶段
FROM alpine:latest

# 安装 ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/server/server-app .

# 暴露 gRPC 端口
EXPOSE 50051

# 运行服务
CMD ["./server-app"]

