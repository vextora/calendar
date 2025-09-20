package main

import (
	"calendarapi/pkg/scaffold"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Generating...")
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/gen/main.go <entity>")
		os.Exit(1)
	}

	version := os.Args[1]
	entity := os.Args[2]
	if err := scaffold.Generate(version, entity); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
