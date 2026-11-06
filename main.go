package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/thuyencode/hello-templ/components"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{Handler: mux, Addr: ":8080"}
	component := components.Hello("World")

	mux.Handle("/", templ.Handler(component))

	log.Fatal(server.ListenAndServe())
}
