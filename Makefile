BIN_NAME := crowned-lang-sv
BIN_EXT := .exe
GIT_TAG := $(shell git describe --abbrev=0 --tags)
GIT_COMMIT := $(shell git rev-list -1 HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME := $(shell date +"%Y-%m-%d %H:%M:%S")
BUILD_DEFINITIONS := -X main.gitTag=$(GIT_TAG) \
-X main.gitBranch=$(GIT_BRANCH) \
-X main.gitCommit=$(GIT_COMMIT) \
-X 'main.buildTime=$(BUILD_TIME)'
ARCHIVE_NAME := crowned-lang-sv-$(GIT_TAG)
RELEASE_ASSETS := $(shell echo $$PWD/release-assets)

SERVER_WINDOWS_AMD64 := windows-amd64/$(BIN_NAME)$(BIN_EXT)
SERVER_LINUX_AMD64 := linux-amd64/$(BIN_NAME)
SERVER_MACOS_AMD64 := macos-amd64/$(BIN_NAME)
SERVER_MACOS_ARM64 := macos-arm64/$(BIN_NAME)

.PHONY: $(SERVER_WINDOWS_AMD64) \
	$(SERVER_LINUX_AMD64) \
	$(SERVER_MACOS_AMD64) \
	$(SERVER_MACOS_ARM64) \
	go-dependencies

# Compilation targets

go-dependencies:
	cd lang/cmd/sv && go mod download

$(SERVER_WINDOWS_AMD64): go-dependencies
	cd lang/cmd/sv && env GOOS=windows GOARCH=amd64 go build -o $@ -ldflags="$(BUILD_DEFINITIONS)"

server-windows-amd64: $(SERVER_WINDOWS_AMD64)
	
$(SERVER_LINUX_AMD64): go-dependencies
	cd lang/cmd/sv && env GOOS=linux GOARCH=amd64 go build -o $@ -ldflags="$(BUILD_DEFINITIONS)"

server-linux-amd64: $(SERVER_LINUX_AMD64)

$(SERVER_MACOS_AMD64): go-dependencies
	cd lang/cmd/sv && env GOOS=darwin GOARCH=amd64 go build -o $@ -ldflags="$(BUILD_DEFINITIONS)"

server-macos-amd64: $(SERVER_MACOS_AMD64)

$(SERVER_MACOS_ARM64): go-dependencies
	cd lang/cmd/sv && env GOOS=darwin GOARCH=arm64 go build -o $@ -ldflags="$(BUILD_DEFINITIONS)"

server-macos-arm64: $(SERVER_MACOS_ARM64)

# Package targets

release-assets:
	mkdir -p release-assets

PACKAGE_WINDOWS_AMD64 := $(RELEASE_ASSETS)/$(ARCHIVE_NAME)-windows-amd64.zip
PACKAGE_LINUX_AMD64 := $(RELEASE_ASSETS)/$(ARCHIVE_NAME)-linux-amd64.tar.gz
PACKAGE_MACOS_AMD64 := $(RELEASE_ASSETS)/$(ARCHIVE_NAME)-macos-amd64.tar.gz
PACKAGE_MACOS_ARM64 := $(RELEASE_ASSETS)/$(ARCHIVE_NAME)-macos-arm64.tar.gz

$(PACKAGE_WINDOWS_AMD64): server-windows-amd64 release-assets
	cd lang/cmd/sv/windows-amd64 && zip $@ -r *

package-windows-amd64: $(PACKAGE_WINDOWS_AMD64)

$(PACKAGE_LINUX_AMD64): server-linux-amd64 release-assets
	cd lang/cmd/sv/linux-amd64 && tar -czf $@ *

package-linux-amd64: $(PACKAGE_LINUX_AMD64)

$(PACKAGE_MACOS_AMD64): server-macos-amd64 release-assets
	cd lang/cmd/sv/macos-amd64 && tar -czf $@ *

package-macos-amd64: $(PACKAGE_MACOS_AMD64)

$(PACKAGE_MACOS_ARM64): server-macos-arm64 release-assets
	cd lang/cmd/sv/macos-arm64 && tar -czf $@ *

package-macos-amd64: $(PACKAGE_MACOS_ARM64)

# Package extension target

package-extension:
	# Update extension version according to tag.
	cd vscode/sv && npm version $(GIT_TAG:v%=%)
	cd vscode/sv && npm ci
	cd vscode/sv && vsce package --out $(RELEASE_ASSETS)

all: package-windows-amd64 package-linux-amd64 package-macos-amd64 package-macos-amd64 package-extension