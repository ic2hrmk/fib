SHELL := /bin/bash

.ONESHELL:
.PHONY: build clean

TARGET_NAME := gib_gen
VERSION = 0.0.1
GITHASH = $(shell git rev-parse --short HEAD)
PROJECT = fib/generator

LDFLAGS += -X $(PROJECT)/main.Version=$(VERSION)-$(GITHASH)

build:
	go build -ldflags "$(LDFLAGS)" -a -o $(TARGET_NAME) main.go