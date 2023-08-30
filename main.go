package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	fmt.Println("Hola mundo!")

	stringPort := os.Getenv("PORT")

	if stringPort == "" {
		log.Fatal("INvalid port :(")
	}

	fmt.Println("Port:", stringPort)
}
