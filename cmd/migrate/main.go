package main

import (
	infra "golang-webserver-practise/internal/infrastructure"
)

func main() {
	infra.Reset()
	// infra.Migrate()
}
