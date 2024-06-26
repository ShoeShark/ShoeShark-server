# 使用官方 Go 镜像作为构建阶段的基础镜像
FROM golang:1.22-alpine AS builder

# 设置 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 复制项目的 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目的源代码
COPY . .

# 编译 Go 程序
RUN go build -o main .

# 使用一个更小的基础镜像来运行编译后的程序
FROM alpine:latest

# 安装 bash
#RUN apk add --no-cache bash

# 设置工作目录
WORKDIR /root/

# 复制编译后的程序
COPY --from=builder /app/main .
COPY --from=builder /app/resources ./resources
COPY --from=builder /app/resources/static ./resources/static

# 暴露端口（根据你的程序的需求来调整）
EXPOSE 8080

# 运行编译后的程序
CMD ["./main"]
