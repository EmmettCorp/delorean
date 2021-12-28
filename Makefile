##@ Test
lint: ## Run linters only.
	@echo -e "\033[2m→ Run linters...\033[0m"
	golangci-lint run --config .golangci.yml

test: ## Run go tests for files with tests.
	@echo -e "\033[2m→ Run tests for all files...\033[0m"
	go test -v ./...

check: lint test ## Run full check: lint and test.

##@ Deploy
build: ## Build binary.
	@echo -e "\033[2m→ Building binary...\033[0m"
	go build -o delorean main.go

install: build ## Install binary to `/usr/local/bin` path.
	@echo -e "\033[2m→ Installing binary to /usr/local/bin ...\033[0m"
	sudo mv delorean /usr/local/bin/
run: ## Run without building
	@echo -e "\033[2m→ Running without building...\033[0m"
	go run main.go

##@ Other
#------------------------------------------------------------------------------
help:  ## Display help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
#------------- <https://suva.sh/posts/well-documented-makefiles> --------------

.DEFAULT_GOAL := help
.PHONY: help lint test check build install run