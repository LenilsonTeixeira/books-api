package main

import (
	"github.com/LenilsonTeixeira/books-api/database"
	"github.com/LenilsonTeixeira/books-api/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()

}
