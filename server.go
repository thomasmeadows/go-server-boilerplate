package main

import (
	"log"
	"go-server-boilerplate/config"
	"net/http"
)


func main() {
	router := config.Router()

	if err := http.ListenAndServe(":3000", &router); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
