# golang http server 学習用

* go 1.20

## run

set environment
```bash
export DB_USER=root
export DB_PASS=
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_NAME=golang-webserver-practise
```

```bash
git clone git@github.com:t1732/golang-webserver-practise.git
cd golang-webserver-practise
go install
docker compose up -d
go run cmd/migrate/main.go
go run cmd/server/main.go
```
