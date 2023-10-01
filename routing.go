package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{name}/page/{subject}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		subject := vars["subject"]

		fmt.Fprintf(w, "You've requested about the user: %s on page %s", name, subject)
	})

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}