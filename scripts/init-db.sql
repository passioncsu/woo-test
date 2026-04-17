-- 学生管理系统 - 数据库初始化脚本
-- Docker 启动时自动执行

-- 创建扩展
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 创建管理员账户（密码: admin123）
-- bcrypt hash of "admin123"
INSERT INTO admins (username, password, created_at, updated_at)
VALUES ('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', NOW(), NOW())
ON CONFLICT (username) DO NOTHING;
