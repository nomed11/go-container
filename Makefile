.PHONY: build clean test setup-rootfs

BINARY_NAME=container
BUILD_DIR=bin

build:
	@echo "Building container runtime..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/container

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@go clean

test:
	@go test ./...

setup-rootfs:
	@echo "Setting up root filesystem..."
	@mkdir -p rootfs
	@if [ ! -f "rootfs/bin/sh" ]; then \
		wget -q https://dl-cdn.alpinelinux.org/alpine/v3.18/releases/x86_64/alpine-minirootfs-3.18.0-x86_64.tar.gz && \
		tar -xzf alpine-minirootfs-3.18.0-x86_64.tar.gz -C rootfs && \
		rm alpine-minirootfs-3.18.0-x86_64.tar.gz; \
	fi