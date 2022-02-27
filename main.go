package main

import (
	"net/http"

	"github.com/afaferz/web-app/routes"
)

func main() {
	routes.HandlerRoutes()
	http.ListenAndServe(":8080", nil)
}
