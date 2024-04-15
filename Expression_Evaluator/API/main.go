package main

import (
	"log"
)

func main() {

	server := NewAPIServer("localhost:8080")

	log.Fatal(server.Run())

}
