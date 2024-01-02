package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "net/http/httputil"
	// "log"
)

type Person struct {
	name string `json:name`
	age int `json:age`
}

func GetJson(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("GetJson starting...\n\n")

	fmt.Printf("%+v\n\n", w)
	fmt.Printf("%+v\n\n", r)
	fmt.Printf("%+v\n\n", r.Body)

	// res, err := httputil.DumpRequest(r, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf(string(res))

	
	// fmt.Printf(r.Body)

	var p Person

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		fmt.Printf("JSON ERROR")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v\n\n", p)


	// fmt.Printf(string(p))
	// fmt.Fprintf(w, "Person: %+v", p)

}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got /hello request\n")
	io.WriteString(w, "Hello, HTTP!")
}
