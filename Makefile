run: build
	@./bin/my-health

.PHONY: diff
include .env
# atlas migrate
migrate:
	atlas schema apply \
          -u '$(DB_URL)?authToken=$(TURSO_TOKEN)'\
          --to file://schema.sql \
					--dev-url "sqlite://dev?mode=memory"
# atlas diff
diff:
	atlas migrate diff $(NAME) \
		  --dir "file://migrations" \
		  --to file://schema.sql \
			--dev-url "sqlite://dev?mode=memory"

#atlas rollback
rollback:
	atlas schema rollback \
		  -u $(DB_URL) \
		  --to file://schema.sql \
			--dev-url "sqlite://dev?mode=memory"


install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

css:
	@tailwindcss -i views/css/app.css -o public/styles.css --watch 

templ:
	@templ generate --watch --proxy=http://localhost:3000

build:
	@templ generate view
	@go build  -o bin/my-health main.go 

