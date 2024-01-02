package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"httpServer/api"
)

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	if p == "./about" {
		p = "./static/about.html"
	}
	http.ServeFile(w, r, p)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {
	http.HandleFunc("/", serveFiles)
	http.HandleFunc("/about", serveFiles)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	// http.HandleFunc("/", getRoot)
	// http.HandleFunc("/hello", api.GetHello)
	http.HandleFunc("/getJson", api.GetJson)

	listenAddr := flag.String("listenaddr", ":3333", "todo")
	flag.Parse()

	fmt.Printf("Starting server at localhost:3333...\n")

	http.ListenAndServe(*listenAddr, nil)

	// fmt.Printf("Checking port is open...\n")
	// l, err := net.Listen("tcp", ":3333")

	// if err != nil {
	// 	fmt.Printf("error listening to port :3333\n%s\n", err)
	// 	os.Exit(1)
	// }
	//
	// fmt.Printf("Serrrver Operrrationnal !\n")
	//
	// log.Fatal(http.Serve(l, nil))
}
