package main

import (
	"log"
	"simple-forum/controller"
)

var port = "3000"

func main()  {
	r := controller.InitRouter()

	// Middleware
	r.InitMiddleware()

	// Routes
	r.Routes()

	// Port
	err := r.RunServer(port)
	if err != nil {
		log.Fatal(err)
	}
}