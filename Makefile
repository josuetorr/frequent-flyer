.PHONY: dev
dev:
	go build -o ./tmp/main ./cmd/server/main.go && air

.PHONY: docker-dev
docker-dev:
	docker compose -f .docker/docker-compose.yml --env-file .dev.env up -d

