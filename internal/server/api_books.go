package server

import (
	"net/http"

	"google.golang.org/protobuf/proto"
)

func (s *Server) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := s.DB.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := proto.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.Write(resp)
}
