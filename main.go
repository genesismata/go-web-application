package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r* http.Request) {
	log.Default(r.RequestURI)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
