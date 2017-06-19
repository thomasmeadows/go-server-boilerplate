package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

func testController (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("This is an test route.\n"))
}

func Test(router *mux.Router) {
	router.HandleFunc("/test", testController).Methods("GET")
}
