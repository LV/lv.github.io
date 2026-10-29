.PHONY: all
all: serve

.PHONY: build
build:
	@go run ./cmd/sitegen
	@echo "site generated successfully"

.PHONY: clean
clean:
	@rm -rf _site
	@echo "_site/ cleaned successfully"

.PHONY: fmt
fmt:
	@go fmt ./...
	@echo "project formatted successfully"

.PHONY: serve
serve: build
	@go run ./cmd/serve
