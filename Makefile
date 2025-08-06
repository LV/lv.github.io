.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  help      - Show this help message"
	@echo "  run       - Run a local Python HTTP server for debugging"
	@echo "  check     - Perform type checks on Javascript code"
	@echo "  feedgen   - Generate feed.atom file"

.PHONY: run
run:
	@echo "Running HTTP Server via python3..."
	@python3 -m http.server -d public

.PHONY: check
check:
	@echo "Checking types via typescript..."
	@tsc --pretty > /dev/null 2>&1 && \
		echo "✅ Type check passed! All is good." || \
		(echo "❌ Type check failed! See errors below:" && tsc)

.PHONY: feedgen
feedgen:
	@echo "Generating feed.atom file..."
	@python3 scripts/feedgen.py
