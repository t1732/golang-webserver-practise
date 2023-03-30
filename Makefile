GO_CMD     := go
GO_RUN     := $(GO_CMD) run
GO_BUILD   := $(GO_CMD) build
GO_TEST    := $(GO_CMD) test -v
GOOS       := $(shell go env GOOS)
# FYI: https://pkg.go.dev/cmd/link
GO_LDFLAGS := -ldflags="-s -w"
AIR_CMD    := air -c .air.toml
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
test:
	env GOOS=$(GOOS) $(GO_TEST) ./...

.PHONEY: db-migrate db-migrate-reset
db-migrate: cmd/migrate/main.go
	@$(GO_RUN) cmd/migrate/main.go
db-migrate-reset: cmd/migrate/main.go
	@$(GO_RUN) cmd/migrate/main.go -m reset

.PHONEY: install-mod echo-linter-install echo-air-install
install-mod:
	@$(GO_CMD) mod tidy
echo-linter-install:
	@echo '\ncheck if golangci-lint is installed.'; \
	if ! type golangci-lint; then \
		echo 'Please install golangci-lint. (brew install golangci-lint)'; \
	fi
echo-air-install:
	@echo '\ncheck if air is installed.'; \
	if ! type air; then \
		echo 'Please install air. (go install github.com/cosmtrek/air@latest)'; \
	fi

BIND_IP = 127.0.0.1
dev-init: install-mod db-migrate-reset echo-linter-install echo-air-install
dev:
	@if type air; then \
		$(MAKE) dev-air; \
	else\
		$(MAKE) dev-run; \
	fi
dev-run: cmd/server/main.go
	@if [ -n "$${PORT}" ]; then \
		$(GO_RUN) cmd/server/main.go -p $${PORT} -b $(BIND_IP); \
	else \
		$(GO_RUN) cmd/server/main.go -b $(BIND_IP); \
	fi
dev-air: cmd/server/main.go
	@if [ -n "$${PORT}" ]; then \
		$(AIR_CMD) -- -p $${PORT} -b $(BIND_IP); \
	else \
		$(AIR_CMD) -- -b $(BIND_IP); \
	fi
