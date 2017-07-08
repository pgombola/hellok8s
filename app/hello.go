package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "Address for server to listen on")
	flag.Parse()

	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello K8s!")
}
