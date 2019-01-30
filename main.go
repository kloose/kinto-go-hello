package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		name := vars["name"]
		fmt.Fprintf(w, "Hello %s\n", name)

	})

	http.ListenAndServe(":8122", r)

}
