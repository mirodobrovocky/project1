package database

type Connection interface {
	FindById(dest interface{}, id interface{}) error
	FindOne(dest interface{}, filter Filter) error
	FindAll(dest interface{}) error
	Save(dest interface{}, save interface{}) error
}
