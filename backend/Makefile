.PHONY: tailwind-watch
tailwind-watch:
	pnpm run tw:watch

.PHONY: tailwind-build
tailwind-build:
	pnpm run tw:build

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: dev
dev:
	go build -o ./tmp/main ./cmd/server/main.go && \
	(make templ-watch & make tailwind-watch &) && \
	air

.PHONY: docker-up
docker-up:
	docker compose -f .docker/docker-compose.yml --env-file .dev.env up -d

