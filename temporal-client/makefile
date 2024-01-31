# Makefile settings
.PHONY: help fmt tidy vet lint stop clean build start run-background test

# Variables
TEMPORAL_DC := ./docker/temporal/docker-compose.yml
DC := docker-compose -f $(TEMPORAL_DC)

# Help target to display available commands
help:
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Format the Go project
fmt: ## Format workspace
	@echo "[fmt] Format go project..."
	@gofmt -l -s -w ./
	@gosimports -l -w ./
	@echo "------------------------------------[Done]"

# Tidy Go modules
tidy: ## Tidy go modules
	@echo "[tidy] Tidying go modules..."
	@rm -f go.sum; go mod tidy -compat=1.21

# Vet Go code
vet: ## Vet go code
	@echo "[vet] Vet go code..."
	@go vet ./...

# Lint Go code
lint: ## Lint go code
	@echo "[lint] Lint go code..."
	@golangci-lint run

# Stop all Docker Containers run in Compose
stop: ## Stop all Docker Containers run in Compose
	@echo "[stop] Stopping Docker containers..."
	@$(DC) stop

# Clean all Docker Containers and Volumes
clean: stop ## Clean all Docker Containers and Volumes
	@echo "[clean] Cleaning up Docker containers and volumes..."
	@$(DC) down --rmi local --remove-orphans -v
	@$(DC) rm -f -v

# Rebuild the Docker Image for use by Compose
build: clean ## Rebuild the Docker Image for use by Compose
	@echo "[build] Building Docker images..."
	@$(DC) build

# Run the Application as a docker compose workflow
start: stop ## Run the Application as a docker compose workflow
	@echo "[start] Starting the application..."
	@$(DC) up

# Run the Application in the background
run-background: stop ## Run the Application in the background
	@echo "[run-background] Running the application in the background..."
	@$(DC) up --detach

# Test the application
test: run-background ## Test the application
	@echo "[test] Running tests..."
	@echo "waiting for services to be ready to accept connections"
	@sleep 60
	@go test ./... -coverprofile cover.out
	@make stop