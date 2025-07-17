.PHONY: all
all: check run

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
	@python3 scripts/feedgen.py
