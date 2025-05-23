# 第一阶段：构建阶段
FROM golang:alpine AS builder

# 设置Go编译环境
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

# 使用清华镜像源加速Alpine包管理
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories

# 更新Alpine包并安装必要工具
RUN apk update --no-cache && \
    apk add --no-cache git tzdata ca-certificates

# 设置工作目录
WORKDIR /build

# 复制Go项目依赖文件
COPY go.mod go.sum ./

# 下载Go依赖
RUN go mod download

# 复制源代码
COPY . /build/

# 编译Go项目
WORKDIR /build/user
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /build/user/user ./user.go

# 第二阶段：运行阶段
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk add --no-cache ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 设置工作目录
WORKDIR /app

# 从构建阶段复制文件并设置执行权限
COPY --from=builder /build/user/user /app/
RUN chmod +x /app/user
COPY --from=builder /build/user/etc/user.yaml /app/etc/user.yaml

# 设置环境变量
ENV TZ=Asia/Shanghai

# 暴露服务端口
EXPOSE 8080

# 运行服务
CMD ["/app/user", "-f", "/app/etc/user.yaml"]