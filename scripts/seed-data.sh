#!/bin/bash
# 填充测试数据
set -e

DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-postgres}"
DB_NAME="${DB_NAME:-student_admin}"

echo ">>> Seeding test data..."

PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME <<'SQL'
-- 插入测试学生数据
INSERT INTO students (name, gender, phone, email, major, class, student_no, status, created_at, updated_at)
VALUES
  ('张三', '男', '13800138001', 'zhangsan@example.com', '计算机科学与技术', '计科2401', '2024001', 1, NOW(), NOW()),
  ('李四', '女', '13800138002', 'lisi@example.com', '软件工程', '软工2401', '2024002', 1, NOW(), NOW()),
  ('王五', '男', '13800138003', 'wangwu@example.com', '数据科学', '数据2401', '2024003', 1, NOW(), NOW()),
  ('赵六', '女', '13800138004', 'zhaoliu@example.com', '人工智能', 'AI2401', '2024004', 2, NOW(), NOW()),
  ('孙七', '男', '13800138005', 'sunqi@example.com', '计算机科学与技术', '计科2402', '2024005', 3, NOW(), NOW()),
  ('周八', '男', '13800138006', 'zhouba@example.com', '信息安全', '安全2401', '2024006', 1, NOW(), NOW()),
  ('吴九', '女', '13800138007', 'wujiu@example.com', '软件工程', '软工2402', '2024007', 1, NOW(), NOW()),
  ('郑十', '男', '13800138008', 'zhengshi@example.com', '物联网工程', '物联2401', '2024008', 1, NOW(), NOW()),
  ('钱十一', '女', '13800138009', 'qian11@example.com', '计算机科学与技术', '计科2401', '2024009', 1, NOW(), NOW()),
  ('陈十二', '男', '13800138010', 'chen12@example.com', '人工智能', 'AI2401', '2024010', 1, NOW(), NOW())
ON CONFLICT (student_no) DO NOTHING;

SELECT COUNT(*) as "学生总数" FROM students;
SQL

echo ">>> Test data seeded successfully!"
