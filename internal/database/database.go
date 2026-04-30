package database

type Database interface {
}

type service struct{}

func New() Database {
	return &service{}
}
