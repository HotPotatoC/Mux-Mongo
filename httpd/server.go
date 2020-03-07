package main

import (
	"log"
	userhandler "mux-mongo/httpd/handler/users"
	"net/http"
)

func main() {
	http.Handle("/", userhandler.Routes())

	log.Fatal(http.ListenAndServe(":8000", nil))
}
