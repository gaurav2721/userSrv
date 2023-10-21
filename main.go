package main

import (
	"os"

	"github.com/gaurav2721/userSrv/database"
	routes "github.com/gaurav2721/userSrv/routes"
	"github.com/gin-gonic/gin"
)

//TODO : pass the listening port in the command line
//TODO : perform basic authentication with hardcoded credential

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" //TODO: check why it has to opt for default port
	}

	//Initilizing database interface , right now we are using in memory to store data , later we can instantiate a noSql client db
	database.DbI = database.InitMemoryDb()

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	router.Run(":" + port)
}
