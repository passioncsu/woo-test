#!/bin/bash
set -e

echo ">>> Deploying to PRODUCTION..."
echo ">>> This is a manual deployment. Please confirm."

COMPOSE_FILE="deploy/docker-compose.yml"

# 数据库备份
echo ">>> Backing up database..."
docker compose -f $COMPOSE_FILE exec -T postgres \
    pg_dump -U postgres student_admin > "backup_$(date +%Y%m%d_%H%M%S).sql"

echo ">>> Pulling latest code..."
git pull origin main

echo ">>> Building and restarting services..."
docker compose -f $COMPOSE_FILE up -d --build --remove-orphans

echo ">>> Running database migrations..."
docker compose -f $COMPOSE_FILE exec -T backend \
    ./server migrate 2>/dev/null || true

echo ">>> Cleaning up..."
docker image prune -f

echo ">>> Production deployment complete!"
docker compose -f $COMPOSE_FILE ps
