.PHONY: all
all: clean compile run

.PHONY: clean
clean:
	rm -rf dist

.PHONY: compile
compile:
	tsc

.PHONY: run
run:
	python3 -m http.server
