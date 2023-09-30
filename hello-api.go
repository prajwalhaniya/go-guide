package main

import (
	"fmt"
	"net/http"
)

func simpleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path);
}

func main() {
	http.HandleFunc("/", simpleApi)

	http.ListenAndServe(":8080", nil)
}