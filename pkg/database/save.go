package database

type BeforeSaveAction interface {
	BeforeSave()
}
