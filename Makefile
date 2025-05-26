.PHONY: all
all: compile server

.PHONY: compile
compile:
	tsc

.PHONY: server
server:
	python3 -m http.server
