# 使用官方 Go 镜像作为基础镜像
FROM golang:latest

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

# 设置工作目录
WORKDIR /root/

# 复制编译后的程序
COPY --from=0 /app/main .

# 暴露端口（根据你的程序的需求来调整）
EXPOSE 8080

# 运行编译后的程序
CMD ["./main"]
