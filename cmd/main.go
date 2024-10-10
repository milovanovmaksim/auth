package main

import (
	"log"

	"github.com/olezhek28/microservices_course_boilerplate/cmd/auth_server"
)

const grpcPort = 50051

func main() {
	server := auth_server.Server{}
	err := server.Start(grpcPort)
	if err != nil {
		log.Fatalf("failed to start a server | err: %v", err)
	}
}
