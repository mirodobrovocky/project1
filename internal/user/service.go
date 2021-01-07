package user

type Service interface {
	GetCurrentUser() (*User, error)
}

type service struct {

}

func (s service) GetCurrentUser() (*User, error) {
	return &User{"CurrentUser"}, nil
}

func NewService() Service {
	return service{}
}
