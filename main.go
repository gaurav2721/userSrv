package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gaurav2721/userSrv/database"
	routes "github.com/gaurav2721/userSrv/routes"
	"github.com/gin-gonic/gin"
)

func GetPort(args []string) string {
	defaultPort := "9000"
	if len(os.Args) > 1 {
		_, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Println("unable to convert port from string to int")
			return defaultPort
		}
		return os.Args[1]
	}
	return defaultPort
}

func main() {
	port := GetPort(os.Args)

	//Initilizing database interface , right now we are using in memory to store data , later we can instantiate a noSql client db
	database.DbI = database.InitMemoryDb()

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	router.Run(":" + port)
}
