package main

import (
	"net/http"
	"log"
)

func main() {
	//initialize Config
	Init_Config()

	//get router
	router := BuildRouter()

	http.Handle("/", router);

	//start listening
	log.Fatal(http.ListenAndServe(":3000", nil))
}