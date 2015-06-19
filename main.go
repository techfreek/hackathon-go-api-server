package main

import (
	"net/http"
	"log"
	"github.com/WSU-ACM/hackathon-go-api-server/Eventbrite"
	"github.com/WSU-ACM/hackathon-go-api-server/Config"
)

func main() {
	//initialize Config
	config.Init_Config()

	//get router
	router := BuildRouter()

	http.Handle("/", router);

	log.Fatal(eventbrite.GetRemainingSpots())

	//start listening
	log.Fatal(http.ListenAndServe(":3000", nil))
}