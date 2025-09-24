package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//  Register handlers
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/artists/", artistHandler)

	fmt.Println("      -      http://localhost:8080      -      ")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
