.PHONY: tailwind-watch
tailwind-watch:
	tailwindcss -i ./assets/style/input.css -o ./assets/style/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	tailwindcss -i ./assets/style/input.css -o ./assets/style/style.min.css --minify

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

