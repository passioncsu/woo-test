#!/bin/bash
set -e

ENV="${1:-dev}"
COMPOSE_OPTS=""

if [ "$ENV" = "prod" ]; then
    COMPOSE_FILE="deploy/docker-compose.yml"
    echo ">>> Deploying to PRODUCTION..."
else
    COMPOSE_FILE="deploy/docker-compose.yml:deploy/docker-compose.dev.yml"
    echo ">>> Deploying to DEVELOPMENT..."
fi

echo ">>> Pulling latest images..."
docker compose -f $COMPOSE_FILE pull

echo ">>> Restarting services..."
docker compose -f $COMPOSE_FILE up -d --remove-orphans

echo ">>> Cleaning up old images..."
docker image prune -f

echo ">>> Deployment complete!"
docker compose -f $COMPOSE_FILE ps
