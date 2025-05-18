# Makefile

export COMPOSE_BAKE = true

ENV_FILE ?= .env
APP_DIR := .
GEN_CMD=./cmd/gen

.DEFAULT_GOAL := help

.PHONY: vendor dev dev-down prod prod-down test lint build run clean coverage coverage-terminal coverage-html gen help

# ------
# Vendor
# ------
vendor:
	@echo ">>> Ensuring go.mod is tidy and vendor is up-to-date"
	cd $(APP_DIR) && go mod tidy && go mod vendor

# -----------
# Development
# -----------
dev: vendor ## Run development environment with Docker
	@echo ">>> Starting development environment"
	@cp .env.development .env.override
	docker compose --env-file .env.override -f docker-compose.dev.yml up --build

dev-down: ## Stop development environment
	@echo ">>> Stopping development environment"
	docker compose --env-file .env.override -f docker-compose.dev.yml down
	@rm -f .env.override

# ----------
# Production
# ----------
prod: ## Run production environment with Docker
	@echo ">>> Starting production environment"
	docker compose --env-file .env -f docker-compose.yml up --build -d

prod-down: ## Stop production environment
	@echo ">>> Stopping production environment"
	docker-compose --env-file .env -f docker-compose.yml down

# -----------------
# Testing & Linting
# -----------------
test: ## Run tests
	@echo ">>> Running tests"
	docker compose --env-file .env.override -f docker-compose.dev.yml exec app go test ./...

lint: ## Run linter
	@echo ">>> Running linter (golangci-lint)"
	docker compose --env-file .env.override -f docker-compose.dev.yml exec app golangci-lint run

# ------------------------
# Build binary static prod
# ------------------------
build: ## Build production binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) ./cmd/server

# ---------------------------------
# Run binary locally (test purpose)
# ---------------------------------
run: ## Run app locally
	go run ./cmd/server/main.go

# ----------------------
# Cleansing local binary
# ----------------------
clean: ## Remove built binaries
	rm -rf $(BINARY_NAME)

# --------
# Coverage
# --------
coverage: ## Run coverage test
	go test ./tests/... -coverpkg=./... -coverprofile=coverage.out

coverage-terminal: ## Show coverage test result at terminal
	go tool cover -func=coverage.out

coverage-html: ## Show coverage test result at browser
	go tool cover -html=coverage.out

# ---------------------------------------------------
# Scaffold entity: make generate ENTITY=[module-name]
# ---------------------------------------------------
gen: ## Generate scaffold for new entity (usage: make gen module-name api-version)
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
	@echo ""
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""
