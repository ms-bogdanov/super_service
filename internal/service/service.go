package service

import (
	"github.com/brianvoe/gofakeit/v7"
	"super_service/internal/repository"
)

type Service struct {
	Repo *repository.UserStorage
}

func NewService(repo *repository.UserStorage) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) TakeBook(userID, bookID int) error {
	err := s.Repo.SearchUser(userID)
	if err != nil {
		return err
	}

	err = s.Repo.SearchBook(bookID)
	if err != nil {
		return err
	}

	err = s.Repo.TakeBook(userID, bookID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ReturnBook(userID, bookID int) error {
	err := s.Repo.SearchUser(userID)
	if err != nil {
		return err
	}

	err = s.Repo.SearchBook(bookID)
	if err != nil {
		return err
	}

	err = s.Repo.ReturnBook(userID, bookID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) InitProject() {
	count := s.Repo.GetAuthorCount()
	if count < 10 {
		s.Repo.AddAuthor(gofakeit.BookAuthor())
	}
}

func (s *Service) InitBook() {
	count := s.Repo.GetBookCount()
	if count < 100 {
		s.Repo.AddBook(gofakeit.BookTitle())
	}
}

func (s *Service) InitUsers() {
	count := s.Repo.GetUsersCount()
	if count < 50 {
		s.Repo.AddUser(gofakeit.Username())
	}
}
