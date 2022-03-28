package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Web server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
