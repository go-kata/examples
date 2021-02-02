L_PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
L_LOCAL_DIR   := $(L_PROJECT_DIR)/.local

.PHONY: all
all:
	@echo "Available commands:"
	@echo ""
	@echo "	example-di-compile		Compile example-di binary"
	@echo ""

.PHONY: example-di.compile
example-di.compile:
	mkdir -p "$(L_LOCAL_DIR)"
	go build -o "$(L_LOCAL_DIR)" "$(L_PROJECT_DIR)/cmd/example-di"
