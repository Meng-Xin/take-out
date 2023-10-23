# 编译阶段：引用最小编译环境
FROM golang:1.21.0 AS builder

# 镜像默认工作目录
WORKDIR /build


# 防止多次拉取依赖
ADD go.mod .
ADD go.sum .
# 配置镜像golang的默认配置,方便拉取依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

# 拷贝当前目录所有文件到工作目录
COPY . .

# 设置编译环境并进行编译
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64  go build -o /app/gin-server .

# 构建阶段：使用 alpine 最小构建
FROM alpine

# 设置镜像工作目录
WORKDIR /app

# 在builder阶段复制可执行的go二进制文件app/go-exporter 到/app/go_exporter中
COPY --from=builder /app/gin-server /app/gin-server
# 创建日志文件夹
RUN mkdir /app/logger

# 时区设置
ENV TZ="Asia/Shanghai"

# 开放端口
EXPOSE 8080

# 启动服务器
CMD ["/app/gin-server"]