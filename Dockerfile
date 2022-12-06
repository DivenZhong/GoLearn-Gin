FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 gin_demo
RUN go build -o gin_demo .

###################
# 移动到用于存放生成的二进制文件的 /dist 目录
###################
WORKDIR /dist

COPY config.yaml .

# 从builder镜像中把/dist/app 拷贝到当前目录
RUN cp /build/gin_demo .


# 需要运行的命令
CMD ["/dist/gin_demo"]