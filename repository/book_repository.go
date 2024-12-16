package repository

import (
	"errors"
	"fmt"
	"handlebooks/db"
	"handlebooks/models"
)

type BookRepository interface {
	GetBook(id int) (*models.Book, error)
	GetAllBooks() ([]models.Book, error)
	CreateBook(book *models.Book) (*models.Book, error)
}

type BookRepositoryDB struct{}

func NewBookRepositoryDB() *BookRepositoryDB {
	return &BookRepositoryDB{}
}

func (r *BookRepositoryDB) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	fmt.Println("error occured1")
	if db.DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	err := db.DB.Find(&books).Error
	fmt.Println("error occured2")
	fmt.Println(books)
	if err != nil {
		fmt.Println("error occured3")
		return nil, err
	}
	fmt.Println(books)
	return books, nil
}

func (r *BookRepositoryDB) CreateBook(book *models.Book) (*models.Book, error) {
	err := db.DB.Create(book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepositoryDB) GetBook(id int) (*models.Book, error) {
	var book models.Book
	err := db.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}
