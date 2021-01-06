package item

type Service interface {
	FindAll() ([]Item, error)
	FindByName(name string) (*Item, error)
	Create(create CreateDto) (*Item, error)
}

type service struct {
	repository Repository
}

func (s service) FindAll() ([]Item, error) {
	return s.repository.FindAll()
}

func (s service) FindByName(name string) (*Item, error) {
	return s.repository.FindByName(name)
}

func (s service) Create(create CreateDto) (*Item, error) {
	user := "CurrentUser" //TODO identification
	return s.repository.Save(Item{
		Name:  create.Name,
		Owner: user,
		Price: create.Price})
}

func NewService(repository Repository) Service {
	return service{repository}
}
