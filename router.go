package main

import(
	"net/http"
	"github.com/WSU-ACM/hackathon-go-api-server/Routes"
)

var Router := http.NewServeMux()
	.Handle("/api/spots", Route.Spots)
	.Handle("/api/teams", Route.Teams)
	.Handle("/api/imgs", Route.Imgs);

