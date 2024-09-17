package service

import (
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
