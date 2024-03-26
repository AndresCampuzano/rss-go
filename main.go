package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot load envs")
		return
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	} else {
		log.Printf("Server is running on port: %v", portString)
	}
}
