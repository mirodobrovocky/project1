package item

import "github.com/mirodobrovocky/project1/internal/user"

type Service interface {
	FindAll() ([]Item, error)
	FindByName(name string) (*Item, error)
	Create(create CreateDto) (*Item, error)
}

type service struct {
	repository 	Repository
	userService user.Service
}

func (s service) FindAll() ([]Item, error) {
	return s.repository.FindAll()
}

func (s service) FindByName(name string) (*Item, error) {
	return s.repository.FindByName(name)
}

func (s service) Create(create CreateDto) (*Item, error) {
	currentUser, err := s.userService.GetCurrentUser()
	if err != nil {
		return nil, err
	}

	return s.repository.Save(NewItem(create.Name, currentUser.Name(), create.Price))
}

func NewService(repository Repository, userService user.Service) Service {
	return service{repository, userService}
}
