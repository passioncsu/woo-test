.PHONY: lint test test-coverage run build clean swagger

# 运行
run:
	go run cmd/server/main.go

# 构建
build:
	go build -o bin/server cmd/server/main.go

# 代码检查
lint:
	golangci-lint run ./...

# 测试
test:
	go test ./... -v

# 测试覆盖率
test-coverage:
	go test ./... -coverprofile=coverage.txt -covermode=atomic
	go tool cover -func=coverage.txt

# Swagger 文档生成
swagger:
	swag init -g cmd/server/main.go -o docs

# 清理
clean:
	rm -rf bin/ coverage.txt

# 下载依赖
deps:
	go mod tidy
