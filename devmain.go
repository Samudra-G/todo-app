package main

import (
	"log"
	"github.com/Samudra-G/todo-app/cmd/api" // replace 'your-module-name' with the actual module name from go.mod
)

func main() {
	log.Println("Launching API server from devmain")
	api.Main() // Expose a Main() in your cmd/api/main.go
}