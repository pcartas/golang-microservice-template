package main

import (
	"github.com/gorilla/mux"
	m "github.com/pcartas/golang-lib/middleware"
	r "github.com/pcartas/golang-lib/router"
)

// generalMiddlewares define los middlewares que se utilizan para todas las rutas
var generalMiddlewares []mux.MiddlewareFunc = []mux.MiddlewareFunc{m.Recovery}

var routes = r.Routes{
	r.Route{
		Name:        "Healthy",
		Method:      "GET",
		Pattern:     "/api/template/healthy",
		HandlerFunc: Healthy,
		Middlewares: nil,
	},
}
