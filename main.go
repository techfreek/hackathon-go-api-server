package main

import (
	"net/http"
	"time"
	"fmt"

	"github.com/WSU-ACM/hackathon-go-api-server/Routes"
)

var server Server

func main() {
	//initialize Config
	Init_Config()

	//set up router
	server := &http.Server{
		Addr: ":3000",
		Handler: Router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//start listening
	err := server.ListenAndServer()
	if err != nil {
		fmt.Println(err)
	}
}