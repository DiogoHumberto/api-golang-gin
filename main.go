package main

import (
	database "github.com/DiogoHumberto/api-go-gin-rest/dataBase"
	"github.com/DiogoHumberto/api-go-gin-rest/routes"
)

func main() {

	database.Connect()

	routes.HandleRequests()
}
