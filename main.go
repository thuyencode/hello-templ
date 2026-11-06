package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/thuyencode/hello-templ/components"
)

func main() {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	port := env["PORT"]
	mux := http.NewServeMux()
	server := http.Server{Handler: mux, Addr: ":" + port}

	mux.Handle("/", templ.Handler(components.Hello("World")))

	log.Fatal(server.ListenAndServe())
}
