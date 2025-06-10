package main

import (
	"log"
	"github.com/Samudra-G/todo-app/cmd/api" 
)

func main() {
	log.Println("Launching API server from devmain")
	api.Main() 
}