// @title Todo App API
// @version 1.0
// @description This is a sample server for a Todo app CLI + API.
// @host localhost:4000
// @BasePath /

package main

import (
	"log"
	"github.com/Samudra-G/todo-app/cmd/api" 
)

func main() {
	log.Println("Launching API server from devmain")
	api.Main() 
}