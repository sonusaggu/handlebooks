package main

import (
	"handlebooks/db"
	"handlebooks/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectDb()

	r := gin.Default()

	r.POST("/graphql", handlers.HandleGraphQLRequests)

	r.Run(":8080")

}
