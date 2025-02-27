.PHONY: tw-build
tw-build:
	pnpm run tw:build

.PHONY: tw-watch
tw-watch:
	pnpm run tw:watch

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: server-dev
server-dev:
	$(MAKE) tw-watch &
	$(MAKE) templ-watch &
	go build -o ./tmp/main ./cmd/server/main.go && air

.PHONY: docker-dev
docker-dev:
	docker compose -f .docker/docker-compose.yml --env-file ./backend/.dev.env up -d

