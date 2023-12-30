package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got /hello request\n")
	io.WriteString(w, "Hello, HTTP!")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	fmt.Printf("Checking port is open...\n")
	// err := http.ListenAndServe(":3333", nil)
	_, err := net.Listen("tcp", ":3333")

	if err != nil {
		fmt.Printf("error listening to port :3333\n%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Serrrver Operrrationnal !\n")

	errServing := http.ListenAndServe(":3333", nil)

	if errors.Is(errServing, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
