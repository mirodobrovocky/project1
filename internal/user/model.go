package user

type User struct {
	name string
}

func (u User) Name() string {
	return u.name
}

func NewUser(name string) User {
	return User{name: name}
}
