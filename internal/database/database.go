package database

import pb "github.com/alcb1310/proto-server/cmd/schemas"

type Database interface {
	GetAuthors() (*pb.Authors, error)
}

type service struct{}

func New() Database {
	return &service{}
}
