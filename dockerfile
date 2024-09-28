FROM golang:1.23 AS builder

ENV GOMODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go mod tidy
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build  -ldflags="-w -s" -o ../main
WORKDIR /app
RUN mkdir publish  \
    && cp main publish  \
    && cp -r config publish

FROM busybox:1.28.4

WORKDIR /app

COPY --from=builder /app/publish .

# 指定运行时环境变量
ENV GIN_MODE=release
EXPOSE 5001

ENTRYPOINT ["./main"]