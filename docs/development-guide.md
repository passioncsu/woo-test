# 开发指南

## Git Flow 分支策略

```
main (生产) ← release/* ← develop (开发主线)
                                      ↑ feature/*
main ← hotfix/*
```

### 分支命名

| 类型 | 命名 | 示例 |
|------|------|------|
| 功能 | feature/xxx | feature/student-export |
| 修复 | fix/xxx | fix/pagination-offset |
| 发布 | release/vX.Y.Z | release/v1.0.0 |
| 紧急修复 | hotfix/xxx | hotfix/login-timeout |

### Commit 规范

使用 [Conventional Commits](https://www.conventionalcommits.org/)：

```
<type>(<scope>): <subject>

feat: 新增学生信息导出功能
fix(student): 修复分页查询偏移量错误
docs: 更新 API 文档
test(auth): 添加登录接口单元测试
refactor(middleware): 重构鉴权中间件
chore(docker): 更新 Docker 镜像版本
```

## 开发工作流

1. **认领任务** — 在 GitLab Issue 中指派给自己
2. **创建分支** — 从 develop 创建 feature/* 分支
3. **本地开发** — 使用 Docker 开发环境
4. **提交代码** — 遵循 Commit 规范，Pre-commit Hook 自动检查
5. **推送分支** — GitLab CI 自动运行 lint + test
6. **创建 MR** — 填写 MR 模板，指派 Reviewer
7. **Code Review** — 至少 1 人通过 + CI 绿灯
8. **合并** — 合并到 develop，自动部署开发环境

## 后端编码规范

### 分层架构

```
Handler (HTTP层) → Service (业务层) → Repository (数据层)
```

- **Handler**：只做参数校验和响应，不包含业务逻辑
- **Service**：核心业务逻辑，可被多个 Handler 复用
- **Repository**：数据库操作，不关心业务

### 命名规范

- 文件名：小写+下划线 `student_repository.go`
- 接口名：大驼峰 `StudentRepository`
- 函数名：大驼峰（导出）/ 小驼峰（私有）
- 常量：大写+下划线 `DEFAULT_PAGE_SIZE`

### 错误处理

```go
// 不要忽略错误
result, err := service.DoSomething()
if err != nil {
    return fmt.Errorf("do something: %w", err)
}
```

## 前端编码规范

### 目录约定

- `src/api/` — 接口定义（按模块分文件）
- `src/views/` — 页面组件（按路由分目录）
- `src/router/routes/modules/` — 路由模块（自动发现）
- `src/components/` — 通用组件

### 新增页面步骤

1. 在 `src/api/modules/` 添加 API
2. 在 `src/views/` 创建页面组件
3. 在 `src/router/routes/modules/` 添加路由
4. 路由自动发现，无需手动注册

## Code Review 检查清单

- [ ] 代码符合编码规范
- [ ] 没有硬编码的密钥/密码
- [ ] 错误处理完善
- [ ] 关键逻辑有注释
- [ ] 新功能有对应测试
- [ ] API 文档已更新

## Merge Request 模板

```markdown
## 变更说明
<!-- 简述本次变更的内容和目的 -->

## 变更类型
- [ ] 新功能 (feat)
- [ ] 修复 (fix)
- [ ] 重构 (refactor)
- [ ] 文档 (docs)
- [ ] 测试 (test)

## 影响范围
<!-- 列出影响的模块/页面 -->

## 测试说明
<!-- 如何验证本次变更 -->

## 截图/录屏
<!-- 如有 UI 变更请附截图 -->
```
