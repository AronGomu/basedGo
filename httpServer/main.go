package main

import (
	"fmt"
	"io"
	"log"
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
	l, err := net.Listen("tcp", ":3333")

	if err != nil {
		fmt.Printf("error listening to port :3333\n%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Serrrver Operrrationnal !\n")

	log.Fatal(http.Serve(l, nil))

	// errServing := http.ListenAndServe(":3333", nil)
	//
	// if errors.Is(errServing, http.ErrServerClosed) {
	// 	fmt.Printf("server closed\n")
	// } else if errServing != nil {
	// 	fmt.Printf("error starting server: %s\n", errServing)
	// 	os.Exit(1)
	// }
	//
	// fmt.Printf("Serrrver Operrrationnal !\n")

}
