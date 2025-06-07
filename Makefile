.PHONY: all
all: run

.PHONY: run
run:
	python3 -m http.server
