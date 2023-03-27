GO_CMD     := go
GO_RUN     := $(GO_CMD) run
GO_BUILD   := $(GO_CMD) build
GO_TEST    := $(GO_CMD) test -v
GOOS       := $(shell go env GOOS)
# FYI: https://pkg.go.dev/cmd/link
GO_LDFLAGS := -ldflags="-s -w"
TARGETS    := bin/server bin/migrate

default: clean build

build: $(TARGETS)
bin/server: cmd/server/main.go
	env GOOS=$(GOOS) $(GO_BUILD) $(GO_LDFLAGS) -o $@ cmd/server/main.go
bin/migrate: cmd/migrate/main.go
	env GOOS=$(GOOS) $(GO_BUILD) $(GO_LDFLAGS) -o $@ cmd/migrate/main.go

.PHONEY: clean fmt lint test
clean:
	rm -rf $(TARGETS) ./vendor
fmt:
	$(GO_CMD) fmt ./...
lint:
	golangci-lint run
	$(GO_CMD) vet ./...
test:
	env GOOS=$(GOOS) $(GO_TEST) ./...

.PHONEY: db-migrate db-migrate-reset
db-migrate: cmd/migrate/main.go
	@$(GO_RUN) cmd/migrate/main.go
db-migrate-reset: cmd/migrate/main.go
	@$(GO_RUN) cmd/migrate/main.go -m reset

.PHONEY: install-mod echo-linter-install
install-mod:
	@$(GO_CMD) mod tidy
echo-linter-install:
	@echo '\ncheck if golangci-lint is installed.'; \
	if ! type golangci-lint; then \
		echo 'Please install golangci-lint. (brew install golangci-lint)'; \
	fi

dev-init: install-mod db-migrate-reset echo-linter-install
dev: cmd/server/main.go
	@if [ -n "$${PORT}" ]; then \
		$(GO_RUN) cmd/server/main.go -p $${PORT}; \
	else \
		$(GO_RUN) cmd/server/main.go; \
	fi
