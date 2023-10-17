package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Check if the endpoint and args are provided
	if len(os.Args) != 4 {
		log.Fatal("Usage: go run main.go <endpoint> <arg1> <arg2>")
		return
	}

	endpoint := os.Args[1]
	arg1 := os.Args[2]
	arg2 := os.Args[3]

	serverURL := "http://localhost:3333/" + endpoint
	switch endpoint {
	case "generate":
		if err := genDockerfile(serverURL, arg1, arg2); err != nil {
			log.Fatalf("Failed to generate Dockerfile: %v", err)
		}
		fmt.Println("Dockerfile has been saved!")
		return
	case "cuegen":
		genK8s(serverURL, arg1, arg2)
	default:

		log.Fatalf("Unsupported endpoint: %s", endpoint)
	}
	fmt.Println("Dockerfile has been saved!")
}
