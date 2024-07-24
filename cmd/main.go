package main

import (
	"log"

	"github.com/berkbeyret/firstgobackendapi/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
