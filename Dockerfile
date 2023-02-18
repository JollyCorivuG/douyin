FROM golang:1.19 as build

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .
RUN G00S=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

FROM scratch as prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/app /

CMD ["/app"]