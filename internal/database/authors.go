package database

import pb "github.com/alcb1310/proto-server/cmd/schemas"

func (s *service) GetAuthors() (pb.Authors, error) {
	a := pb.Authors{
		Authors: []*pb.Author{
			{
				Id:   1,
				Name: "Martin Kleppmann",
			},
			{
				Id:   2,
				Name: "Alex Petrov",
			},
			{
				Id:   3,
				Name: "Isaac Asimov",
			},
		},
	}

	return a, nil
}
