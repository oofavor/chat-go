package handlers

import (
	"fmt"
	"net/http"
)

func Example(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "yes")
}
