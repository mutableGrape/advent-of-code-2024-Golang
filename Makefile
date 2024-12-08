# Variables
BINARY_NAME=aoc
BINARY_PATH=bin/$(BINARY_NAME)

# Default target
all: build

build: 
	go build -o $(BINARY_PATH) ./cli

run: 
	build $(BINARY_PATH) $(ARGS)

clean: 
	rm -f $(BINARY_PATH)

.PHONY: all build run clean
