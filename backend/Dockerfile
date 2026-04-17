# 多阶段构建
# 阶段1: 构建
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /bin/server cmd/server/main.go

# 阶段2: 运行
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata musl
ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder /bin/server .
COPY config.yaml .

EXPOSE 8080

CMD ["./server"]
