package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

func rootController (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("This is an example server.\n"))
}

func Root(router *mux.Router) {
	router.HandleFunc("/", rootController).Methods("GET")
}
