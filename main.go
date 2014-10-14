package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	optFile := flag.String("f", "", "Path to serve")
	optPort := flag.String("p", ":8080", "Port to serve")
	flag.Parse()
	file := *optFile
	port := *optPort
	if file == "" {
		log.Fatal("Nothing to serve.")
	}
	http.Handle("/", http.FileServer(http.Dir(file)))
	log.Println("Serving", file, "on", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
