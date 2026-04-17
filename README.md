# 学生管理系统 (Student Admin)

基于 Vue Vben Admin + Go + PostgreSQL 的学生信息管理系统。

## 技术栈

| 层 | 技术 |
|---|------|
| 前端 | Vue 3 + Vben Admin v5 + Ant Design Vue + TypeScript + Vite |
| 后端 | Go 1.22 + Gin + GORM |
| 数据库 | PostgreSQL 15 |
| 缓存 | Redis 7 |
| 认证 | JWT (Bearer Token) |
| 部署 | Docker Compose + Nginx |

## 快速启动

### 前置要求

- Go 1.22+
- Node.js 20.15+
- pnpm (通过 corepack 启用)
- Docker & Docker Compose
- PostgreSQL 15

### 1. 启动开发环境

```bash
# 启动 PostgreSQL + Redis
cd deploy
cp .env.example .env
docker compose -f docker-compose.yml up -d postgres redis

# 启动后端
cd ../backend
cp config.yaml config.local.yaml  # 按需修改数据库连接
make run

# 启动前端
cd ../frontend
corepack enable
pnpm install
pnpm run dev:antd
```

### 2. 访问系统

- 前端：http://localhost:5666
- 后端 API：http://localhost:8080/api/

### 3. 初始化数据

```bash
# 创建管理员账户（首次运行后端会自动迁移表结构）
cd backend && go run cmd/server/main.go

# 或使用 seed 脚本填充测试数据
bash scripts/seed-data.sh
```

## 项目结构

```
student-admin/
├── backend/                 # Go 后端
│   ├── cmd/server/          # 入口
│   ├── internal/            # 内部包（分层架构）
│   │   ├── config/          # 配置管理
│   │   ├── handler/         # HTTP 处理器
│   │   ├── service/         # 业务逻辑
│   │   ├── repository/      # 数据访问
│   │   ├── middleware/       # JWT/CORS/日志
│   │   ├── model/           # 数据模型
│   │   └── router/          # 路由注册
│   ├── pkg/                 # 可复用工具包
│   ├── tests/               # 集成测试
│   ├── Dockerfile
│   └── Makefile
├── frontend/                # Vben Admin v5 前端
│   └── apps/web-antd/       # Ant Design Vue 版本
│       └── src/
│           ├── api/         # 接口封装
│           ├── views/       # 页面
│           ├── router/      # 路由
│           └── store/       # 状态管理
├── deploy/                  # 部署配置
│   ├── docker-compose.yml
│   ├── docker-compose.dev.yml
│   ├── nginx/nginx.conf
│   └── .env.example
├── scripts/                 # 辅助脚本
├── docs/                    # 文档
└── .gitlab-ci.yml           # CI/CD 流水线
```

## API 概览

| 方法 | 路径 | 说明 | 鉴权 |
|------|------|------|------|
| POST | /api/login | 管理员登录 | 否 |
| POST | /api/register | 注册管理员 | 否 |
| GET | /api/profile | 获取当前用户信息 | 是 |
| GET | /api/students | 学生列表（分页+搜索） | 是 |
| POST | /api/students | 新增学生 | 是 |
| GET | /api/students/:id | 学生详情 | 是 |
| PUT | /api/students/:id | 更新学生 | 是 |
| DELETE | /api/students/:id | 删除学生 | 是 |

## 常用命令

```bash
# 后端
cd backend
make run            # 运行
make test           # 测试
make test-coverage  # 测试覆盖率
make lint           # 代码检查
make build          # 构建

# 前端
cd frontend
pnpm run dev:antd   # 开发
pnpm run build:antd # 构建
pnpm run lint       # 代码检查
```

## 生产部署

```bash
cd deploy
cp .env.example .env
# 修改 .env 中的密码和密钥
docker compose up -d
```

## 开发流程

详见 [开发指南](docs/development-guide.md)
