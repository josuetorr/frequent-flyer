.PHONY: dev
dev:
	cd backend && go build -o ./tmp/main ./cmd/server/main.go && air

.PHONY: docker-dev
docker-dev:
	docker compose -f .docker/docker-compose.yml --env-file ./backend/.dev.env up -d

