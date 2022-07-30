package main

import (
	"github.com/marcellorhcp/webapi-mvc/database"
	"github.com/marcellorhcp/webapi-mvc/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
