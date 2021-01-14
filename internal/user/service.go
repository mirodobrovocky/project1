package user

type Service interface {
	GetCurrentUser() (*User, error)
}

type service struct {
	repository Repository
}

func (s service) GetCurrentUser() (*User, error) {
	return s.repository.FindById(1)
}

func NewService(repository Repository) Service {
	return service{repository: repository}
}
