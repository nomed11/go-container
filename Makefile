.PHONY: build clean test setup-rootfs

BINARY_NAME=container
BUILD_DIR=bin

# Detect architecture
ARCH := $(shell uname -m)
ALPINE_VERSION=3.18.0

build:
	@echo "Building container runtime..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/container

clean:
	@echo "Cleaning up..."
	@sudo umount -l rootfs/proc 2>/dev/null || true
	@sudo umount -l rootfs/sys 2>/dev/null || true
	@sudo umount -l rootfs/dev 2>/dev/null || true
	@sudo umount -l rootfs/oldrootfs 2>/dev/null || true
	@sudo umount -l rootfs 2>/dev/null || true
	@rm -rf $(BUILD_DIR)
	@sudo rm -rf rootfs/*
	@go clean

test:
	@go test ./...

setup-rootfs:
	@echo "Setting up root filesystem..."
	@sudo mkdir -p rootfs
	@if [ ! -f "rootfs/bin/sh" ]; then \
		if [ "$(ARCH)" = "aarch64" ]; then \
			echo "Detected ARM64 architecture, downloading ARM64 Alpine Linux..." && \
			wget -q https://dl-cdn.alpinelinux.org/alpine/v3.18/releases/aarch64/alpine-minirootfs-$(ALPINE_VERSION)-aarch64.tar.gz && \
			sudo tar -xzf alpine-minirootfs-$(ALPINE_VERSION)-aarch64.tar.gz -C rootfs --no-same-owner && \
			rm alpine-minirootfs-$(ALPINE_VERSION)-aarch64.tar.gz; \
		else \
			echo "Detected x86_64 architecture, downloading x86_64 Alpine Linux..." && \
			wget -q https://dl-cdn.alpinelinux.org/alpine/v3.18/releases/x86_64/alpine-minirootfs-$(ALPINE_VERSION)-x86_64.tar.gz && \
			sudo tar -xzf alpine-minirootfs-$(ALPINE_VERSION)-x86_64.tar.gz -C rootfs --no-same-owner && \
			rm alpine-minirootfs-$(ALPINE_VERSION)-x86_64.tar.gz; \
		fi \
	fi
	@echo "Root filesystem setup complete for $(ARCH)"