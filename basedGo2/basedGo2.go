package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AronGomu/basedGo/basedGo"
)

func main() {

	fmt.Println(os.Args)

	log.SetPrefix("basedGo2: ")
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatal("basedGo2 has no command line argument !")
	}

	message, err := basedGo.Hello(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
