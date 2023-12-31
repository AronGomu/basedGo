package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetJson(w http.ResponseWriter, r *http.Request) {
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got /hello request\n")
	io.WriteString(w, "Hello, HTTP!")
}
