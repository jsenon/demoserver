//go:generate swagger generate spec
// Package main demoserver.
//
// the purpose of this application is to provide an CMDB application
// that will store information in mongodb backend
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
package main

import (
	"api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"web"
)

func main() {

	r := mux.NewRouter()

	// Remove CORS Header check to allow swagger and application on same host and port
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	// To be changed
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "PATCH"})

	// Web Part
	r.HandleFunc("/", web.Index)

	// Swagger
	// r.HandleFunc("/api", web.ApiHelp)

	// API Part

	// Health Check
	r.HandleFunc("/healthy/am-i-up", api.Statusamiup).Methods("GET")
	r.HandleFunc("/healthy/about", api.Statusabout).Methods("GET")

	http.ListenAndServe(":9010", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
