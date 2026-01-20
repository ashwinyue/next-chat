FROM golang:1.24-alpine AS builder

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/aicenter ./cmd/server

# 运行镜像
FROM alpine:3.19

WORKDIR /app

# 安装 ca-certificates（用于 HTTPS 请求）
RUN apk add --no-cache ca-certificates tzdata

# 复制二进制文件
COPY --from=builder /app/bin/aicenter /app/aicenter

# 复制配置文件
COPY --from=builder /app/configs /app/configs

# 暴露端口
EXPOSE 8081

# 运行
CMD ["/app/aicenter"]
