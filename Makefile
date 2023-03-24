GO_CMD     := go
GO_BUILD   := $(GO_CMD) build
GO_TEST    := $(GO_CMD) test -v
GO_VET     := $(GO_CMD) vet
GO_FMT     := $(GO_CMD) fmt
GO_LDFLAGS := -ldflags="-s -w" # FYI: https://pkg.go.dev/cmd/link
GOOS       := $(shell go env GOOS)
TARGETS    := bin/server

.PHONEY: default build clean test fmt lint

default: fmt vet test build
build: $(TARGETS)

bin/server:
	env GOOS=$(GOOS) $(GO_BUILD) $(GO_LDFLAGS) -o $@ cmd/server/main.go
clean:
	rm -rf $(TARGETS) ./vendor
vet:
	$(GO_VET) ./...
test:
	env GOOS=$(GOOS) $(GO_TEST) ./...
fmt:
	$(GO_FMT) ./...
