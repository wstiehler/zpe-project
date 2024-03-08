
include Makefile.coverage.mk
include Makefile.tests.mk
include Makefile.nats.mk

PROJECT_NAME := "zpe-api"


.PHONY: all dep build clean lint

setup: ## Setup project
	@echo "Get the dependencies..."
	@make dep --silent 
	@echo "Install staticcheck to lint..."
	@go install honnef.co/go/tools/cmd/staticcheck@2022.1.2
	@echo "Install gosec to lint..."
	@go install github.com/securego/gosec/v2/cmd/gosec@v2.13.1
	@echo "Configuring hooks..."
	@git config core.hooksPath hooks/
	@chmod +x ./hooks/pre-commit
	@echo "Done."

all: build

dep-dev-run: ## Run development dependencies
	@docker-compose -f docker-compose.yml up -d  --build --remove-orphans

dep-dev-stop: ## Stop development dependencies
	@docker-compose stop

start-ngrok:
	chmod +x config/scripts/start-ngrok.sh && config/scripts/start-ngrok.sh

dep-dev-status: ## Show status from development dependencies
	@docker-compose status

dev-start-with-tools: dep-dev-run ## Run application and dependencies
	@docker-compose -f docker-compose.app.yml -f docker-compose.tools.yml -f docker-compose.yml up -d  --build --remove-orphans

dev-start-with-db: dep-dev-run ## Run application and dependencies
	@docker-compose -f  docker-compose.yml -f docker-compose.app.yml up -d  --build --remove-orphans

dev-down: 
	@docker-compose -f  docker-compose.yml -f docker-compose.app.yml down -v

dev-start-db: dep-dev-run ## Run application and dependencies
	@docker-compose up -d  --build --remove-orphans

show-logs: ## Show logs from development dependencies
	@docker logs ngps-api -f --tail=100

create-network-external:
	@docker network create openvagas

create-network:
	@docker network create openvagas

lint: ## Lint the files
	@staticcheck ./...

dep: ## Get the dependencies
	@go get -v -d ./...

build: dep ## Build the binary file
	@go build -v -o bin/${PROJECT_NAME} web/main.go

clean: ## Remove previous build
	@rm -f bin/$(PROJECT_NAME)

help: ## Show commands and description
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

security: ## Check security vulnerabilities
	@govulncheck ./...