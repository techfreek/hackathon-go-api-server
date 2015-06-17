package main

import(
	"github.com/gorilla/mux"
	"github.com/WSU-ACM/hackathon-go-api-server/Routes"
)

func BuildRouter() *mux.Router {
	m := mux.NewRouter().StrictSlash(false)
	api := m.Path("/api").Subrouter();

	//create routes
	api.Methods("GET").Path("/teams").HandlerFunc(Route.Teams);
	api.Methods("GET").Path("/spots").HandlerFunc(Route.Spots);
	api.Methods("GET").Path("/imgs/{year}").HandlerFunc(Route.Imgs);
	//May need the following route in the case of no year given 
	//api.Methods("GET").Path("/imgs").HandlerFunc(Route.Imgs);

	return api
}