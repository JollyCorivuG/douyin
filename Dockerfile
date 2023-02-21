FROM golang:1.19

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=goproxy.io \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /project/douyin

# 复制go.mod，go.sum并且下载依赖
COPY go.* ./
RUN go mod download

RUN apt update && apt install -y ffmpeg

# 复制项目内的所有内容并构建
COPY . .
RUN go build -o /project/douyin/build/myapp .

EXPOSE 8080
ENTRYPOINT ["/project/douyin/build/myapp"]