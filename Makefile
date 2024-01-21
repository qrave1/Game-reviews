GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go install github.com/google/wire/cmd/wire@latest


.PHONY: run
run:
	@go run backend/cmd/app/main.go


.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: up
up:
	@docker-compose up -d

.PHONY: down
down:
	@docker-compose down
