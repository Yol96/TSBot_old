.PHONY: build
build:
	go build -v ./cmd/tsbot

.DEFAULT_GOAL := build