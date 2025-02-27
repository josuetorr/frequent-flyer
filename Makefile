.PHONY: web-dev
web-dev:
	cd web && pnpm run dev

.PHONY: server-dev
server-dev:
	cd backend && go build -o ./tmp/main ./cmd/server/main.go && air

.PHONY: docker-dev
docker-dev:
	docker compose -f .docker/docker-compose.yml --env-file ./backend/.dev.env up -d

