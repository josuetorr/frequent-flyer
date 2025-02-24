#!/bin/sh

docker compose -f .docker/docker-compose.yml --env-file .dev.env up -d
