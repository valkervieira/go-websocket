package main

import (
	"fmt"
	"net/http"
) 

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Web socket")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndPoint)
}

func main() {

	setupRoutes()
	http.ListenAndServe(":8080", nil)
}