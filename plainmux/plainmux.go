package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/halo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("application/type", "JSON")
		data := "Ini via mux"
		json, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	})

	http.ListenAndServe(":8000", r)

}
