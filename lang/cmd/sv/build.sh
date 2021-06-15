#!/bin/bash
BIN_NAME="crowned-lang-sv"
BIN_EXT=".exe"
GIT_TAG=$(git describe --abbrev=0 --tags)
GIT_COMMIT=$(git rev-list -1 HEAD)
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
BUILD_TIME=$(date +"%Y-%m-%d %H:%M:%S")

# GOOS: windows, linux, darwin. GOARCH: amd64, arm64
GOOS=windows GOARCH=amd64 go build -o $BIN_NAME$BIN_EXT -ldflags \
  "-X main.gitTag=$GIT_TAG       \
   -X main.gitBranch=$GIT_BRANCH \
   -X main.gitCommit=$GIT_COMMIT \
   -X main.buildTime=$BUILD_TIME"
