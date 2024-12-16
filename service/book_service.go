package service

import (
	"handlebooks/models"
	"handlebooks/repository"
)

// BookService defines methods for business logic related to books.
type BookService struct {
	Repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return BookService{Repo: repo}
}

func (s *BookService) CreateBook(book models.Book) (*models.Book, error) {
	return s.Repo.CreateBook(&book)
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.Repo.GetAllBooks()
}

func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.Repo.GetBook(id)
}

// func (s *BookService) UpdateBook(id int, title, author string) (*models.Book, error) {
// 	book, err := s.Repo.GetByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	book.Title = title
// 	book.Author = author
// 	return s.Repo.Update(book)
// }

// func (s *BookService) DeleteBook(id int) error {
// 	return s.Repo.Delete(id)
// }
