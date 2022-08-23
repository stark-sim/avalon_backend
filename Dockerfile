FROM --platform=$BUILDPLATFORM golang:1.18-alpine AS builder

LABEL maintainer="StarkSim<gooda159753@163.com>"

# 在容器根目录创建 src 目录
WORKDIR /src

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache g++

COPY ./go.mod .

COPY ./go.sum .

ENV GOPROXY="https://goproxy.cn"

RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLE=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags "-s -w" -o apiserver ./

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /src/apiserver /app/
COPY --from=builder /src/migrations /app/migrations/

EXPOSE 8081

ENTRYPOINT ["./apiserver"]
