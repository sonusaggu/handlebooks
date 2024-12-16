package graphql

import (
	"fmt"
	"handlebooks/models"
	"handlebooks/service"

	"github.com/graphql-go/graphql"
)

var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.Int},
		"title":  &graphql.Field{Type: graphql.String},
		"author": &graphql.Field{Type: graphql.String},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"book": &graphql.Field{
			Type:        bookType,
			Description: "Get a book by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if !ok {
					return nil, fmt.Errorf("id must be an integer")
				}
				bookService := p.Context.Value("bookService").(service.BookService)
				book, err := bookService.GetBookByID(id)

				if err != nil {
					return nil, err
				}

				return book, nil
			},
		},
		"books": &graphql.Field{
			Type:        graphql.NewList(bookType),
			Description: "Get all books",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				bookService := p.Context.Value("bookService").(service.BookService)
				books, err := bookService.GetAllBooks()

				if err != nil {
					return nil, err
				}

				return books, nil
			},
		},
	},
})

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createBook": &graphql.Field{
				Type:        bookType,
				Description: "Create a new book",
				Args: graphql.FieldConfigArgument{
					"title":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"author": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					bookService := p.Context.Value("bookService").(service.BookService)
					book := models.Book{
						Title:  p.Args["title"].(string),
						Author: p.Args["author"].(string),
					}
					_, err := bookService.CreateBook(book)
					if err != nil {
						// If there is an error, return the error
						return nil, err
					}
					return book, nil
				},
			},
		},
	})

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
