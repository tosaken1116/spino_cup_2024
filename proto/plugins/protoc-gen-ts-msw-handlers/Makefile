# Variables
PROTO_DIR := ./sample
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)
PROTO_OUT_DIR := .
PROTO_PLUGIN_NAME := ts-msw-handlers
PROTO_PLUGIN_FILE := protoc-gen-$(PROTO_PLUGIN_NAME)

# Targets
.PHONY: all clean


build:
	@go build -o $(PROTO_PLUGIN_FILE) .

clean:
	@rm -rf pb
	@rm -f $(PROTO_PLUGIN_FILE)