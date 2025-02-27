.PHONE: tw-build
tw-build:
	pnpm run tw:build

.PHONE: tw-watch
tw-watch:
	pnpm run tw:watch

.PHONY: server-dev
server-dev:
	go build -o ./tmp/main ./cmd/server/main.go && air

.PHONY: docker-dev
docker-dev:
	docker compose -f .docker/docker-compose.yml --env-file ./backend/.dev.env up -d

