package handlers

import (
	gql "handlebooks/graphql"
	"handlebooks/repository"
	"handlebooks/service"
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/gin-gonic/gin"
)

func HandleGraphQLRequests(c *gin.Context) {
	var params struct {
		Query string
	}

	bookRepo := repository.NewBookRepositoryDB()
	bookService := service.NewBookService(bookRepo)

	// Pass bookService into the context for use in resolvers
	c.Set("bookService", bookService)

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	// Execute the query using the schema
	result := graphql.Do(graphql.Params{
		Schema:        gql.Schema,   // Pass the schema here
		RequestString: params.Query, // The query string
		Context:       c,
	})

	if len(result.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": result.Errors})
		return
	}
	c.JSON(http.StatusOK, result.Data)
}
