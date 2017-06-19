package config

import (
	"github.com/gorilla/mux"
	"go-server-boilerplate/routes"
)

func Router() mux.Router {

	Router := mux.NewRouter()

	routes.Root(Router);
	routes.Test(Router);

	return *Router;
}
