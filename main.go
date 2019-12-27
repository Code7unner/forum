package main

import (
	"log"
	"simple-forum/controller"
	db2 "simple-forum/db"
)

var port = "3000"

func main()  {
	db, err := db2.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	r := controller.InitController(db.DB)

	// Middleware
	r.InitMiddleware()

	// Routes
	r.Routes()

	// Port
	err = r.RunServer(port)
	if err != nil {
		log.Fatal(err)
	}
}