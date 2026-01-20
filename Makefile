.PHONY: all build run test clean docker-up docker-down

# 默认目标
all: build

# 构建所有服务
build:
	@echo "Building aicenter..."
	cd aicenter && go build -o bin/aicenter ./cmd/server

# 运行 aicenter
run-aicenter:
	cd aicenter && go run ./cmd/server

# 测试
test:
	cd aicenter && go test ./...

# 清理
clean:
	rm -rf aicenter/bin

# Docker 操作
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

# 数据库迁移
migrate-up:
	cd aicenter && go run ./cmd/migrate up

migrate-down:
	cd aicenter && go run ./cmd/migrate down

# 开发环境（仅启动依赖）
dev-deps:
	docker-compose up -d postgres redis

# 格式化
fmt:
	cd aicenter && go fmt ./...

# Lint
lint:
	cd aicenter && golangci-lint run
