# golang http server 学習用

* go 1.20

## run

set environment

```bash
export DB_USER=root
export DB_PASS=
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_NAME=golang_webserver_practise_development
```

```bash
git clone git@github.com:t1732/golang-webserver-practise.git
cd golang-webserver-practise
docker compose up -d
make dev-init
make dev
```

サーバの待受 port を 8080 に変更したい場合 (デフォルト 3000 ポート)

```bash
make dev PORT=8080
```

## live reload

```bash
go install github.com/cosmtrek/air@latest
make dev
```

## linter

[golangci-lint](https://golangci-lint.run/)

install

```bash
brew install golangci-lint
```

run

```bash
make lint
```
