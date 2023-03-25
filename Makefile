GO_CMD     := go
GO_RUN     := $(GO_CMD) run
GO_BUILD   := $(GO_CMD) build
GO_TEST    := $(GO_CMD) test -v
GO_VET     := $(GO_CMD) vet
GO_FMT     := $(GO_CMD) fmt
GO_LDFLAGS := -ldflags="-s -w" # FYI: https://pkg.go.dev/cmd/link
GOOS       := $(shell go env GOOS)
TARGETS    := bin/server bin/migrate

ifeq ($(PORT),)
  RUN_SERVER := $(GO_RUN) cmd/server/main.go
else
  RUN_SERVER := $(GO_RUN) cmd/server/main.go -p $(PORT)
endif

.PHONEY: default build clean test fmt lint

default: fmt vet test build server db/migrate db/reset
build: $(TARGETS)
s: server

bin/server:
	env GOOS=$(GOOS) $(GO_BUILD) $(GO_LDFLAGS) -o $@ cmd/server/main.go
bin/migrate:
	env GOOS=$(GOOS) $(GO_BUILD) $(GO_LDFLAGS) -o $@ cmd/migrate/main.go
clean:
	rm -rf $(TARGETS) ./vendor
vet:
	$(GO_VET) ./...
test:
	env GOOS=$(GOOS) $(GO_TEST) ./...
fmt:
	$(GO_FMT) ./...
server:
	$(RUN_SERVER)
db/migrate:
	$(GO_RUN) cmd/migrate/main.go
db/reset:
	$(GO_RUN) cmd/migrate/main.go -m reset
