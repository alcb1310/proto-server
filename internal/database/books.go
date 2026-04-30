package database

import pb "github.com/alcb1310/proto-server/cmd/schemas"

func (s *service) GetBooks() (*pb.Books, error) {
	books := &pb.Books{
		Books: []*pb.Book{
			{
				Id:    1,
				Title: "The Lord of the Rings",
				Authors: []*pb.Author{
					{
						Id:   1,
						Name: "J. R. R. Tolkien",
					},
				},
				Publisher: &pb.Publisher{
					Id:   1,
					Name: "Allen & Unwin",
				},
				Genre: pb.Genre_FANTASY,
			},
			{
				Id:    2,
				Title: "The C programming language",
				Authors: []*pb.Author{
					{
						Id:   2,
						Name: "Brian Kernighan",
					},
					{
						Id:   3,
						Name: "Dennis M. Ritchie",
					},
				},
				Publisher: &pb.Publisher{
					Id:   2,
					Name: "Prentice Hall",
				},
				Genre: pb.Genre_TECHNICAL,
			},
		},
	}

	return books, nil
}
