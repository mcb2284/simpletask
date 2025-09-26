package main

import (
	"github.com/simpletask/api"
	"github.com/simpletask/database"
)

func main() {

	database.Database()

	api.InitServer()

}
