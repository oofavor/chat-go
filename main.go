package main

import (
	"log"
	"net/http"

	"github.com/oofavor/go-chat/handlers"
)

func main() {
    http.HandleFunc("/example", handlers.Example)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
