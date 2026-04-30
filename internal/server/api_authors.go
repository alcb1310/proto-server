package server

import (
	"net/http"

	"google.golang.org/protobuf/proto"
)

func (s *Server) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := s.DB.GetAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := proto.Marshal(&authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.Write(resp)
}
