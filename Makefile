# Makefile

export COMPOSE_BAKE = true

ENV_FILE ?= .env
APP_DIR := .
GEN_CMD=./cmd/gen

# Color codes
RED     := \033[1;31m
YELLOW  := \033[1;33m
BLUE    := \033[1;34m
CYAN    := \033[1;36m
GREEN   := \033[1;32m
MAGENTA := \033[1;35m
RESET   := \033[0m

.DEFAULT_GOAL := help

.PHONY: vendor dev dev-down prod prod-down test lint build run clean coverage coverage-terminal coverage-html gen help

vendor: ## Update `go.mod` and vendor
	@echo ">>> Ensuring go.mod is tidy and vendor is up-to-date"
	cd $(APP_DIR) && go mod tidy && go mod vendor

# --- Target kategori Dev ---
dev: vendor ## Run development environment with Docker
	@echo ">>> Starting development environment"
	docker compose -f docker-compose.dev.yml up --build

dev-down: ## Stop development environment
	@echo ">>> Stopping development environment"
	docker compose -f docker-compose.dev.yml down

# --- Target kategori Prod ---
prod: ## Run production environment with Docker
	@echo ">>> Starting production environment"
	docker compose --env-file .env -f docker-compose.yml up --build -d

prod-down: ## Stop production environment
	@echo ">>> Stopping production environment"
	docker-compose --env-file .env -f docker-compose.yml down

build: ## Build binary static for production
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) ./cmd/server

lint: ## Run linter
	@echo ">>> Running linter (golangci-lint)"
	docker compose -f docker-compose.dev.yml exec app golangci-lint run

run: ## Run binary locally (test purpose)
	go run ./cmd/server/main.go

clean: ## Remove built binaries
	rm -rf $(BINARY_NAME)

# --- Target kategori Test ---
test: ## Run tests
	@echo ">>> Running tests"
	docker compose -f docker-compose.dev.yml exec app go test ./...

# --- Target kategori Cov ---
coverage: ## Run coverage test
	go test ./tests/... -coverpkg=./... -coverprofile=coverage.out

coverage-terminal: ## Show coverage test result at terminal
	go tool cover -func=coverage.out

coverage-html: ## Show coverage test result at browser
	go tool cover -html=coverage.out

# --- Target kategori Gen ---
gen: ## Generate scaffold for new entity (usage: make gen <version> <module-name>)
	@if [ $(words $(MAKECMDGOALS)) -ne 3 ]; then \
		echo "Usage: make gen <version> <entity>"; \
		exit 1; \
	fi; \
	go run $(GEN_CMD) $(word 2, $(MAKECMDGOALS)) $(word 3, $(MAKECMDGOALS))

.PHONY: $(wordlist 2, 99, $(MAKECMDGOALS))
$(wordlist 2, 99, $(MAKECMDGOALS)):
	@:

%:
	@echo "\033[31mUnknown command:\033[0m '$@'"
	@$(MAKE) help

help:
	@echo "$(MAGENTA)Usage: make <target>$(RESET)\n"
	@echo "$(RED)Available commands:$(RESET)"

	@echo "$(CYAN)Development:$(RESET)"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep '^dev' | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(RESET) %s\n", $$1, $$2}'

	@echo "$(CYAN)Production:$(RESET)"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep '^prod' | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(RESET) %s\n", $$1, $$2}'

	@echo "$(CYAN)Test:$(RESET)"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep '^test' | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(RESET) %s\n", $$1, $$2}'

	@echo "$(CYAN)Coverage:$(RESET)"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep '^cov' | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(RESET) %s\n", $$1, $$2}'

	@echo "$(CYAN)Generate:$(RESET)"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep '^gen' | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(RESET) %s\n", $$1, $$2}'

	@echo "$(CYAN)Others:$(RESET)"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		grep -v -E '^(prod|dev|test|cov|gen)' | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(RESET) %s\n", $$1, $$2}'
	@echo ""
