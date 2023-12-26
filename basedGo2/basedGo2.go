package main

import (
	"fmt"
	"log"

	"github.com/AronGomu/basedGo/basedGo"
)

func main() {

	log.SetPrefix("basedGo2: ")
	log.SetFlags(0)

	message, err := basedGo.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
