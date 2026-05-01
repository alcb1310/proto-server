package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func (s *Server) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := s.DB.GetAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := proto.Marshal(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("% x\n", resp)

	w.Header().Set("Content-Type", "application/protobuf")
	w.Write(resp)
}

func (s *Server) JSONGetAllAuthors(w http.ResponseWriter, r *http.Request) {
	encodedAuthors, err := s.DB.GetAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authors := make([]map[string]any, 0)
	for _, author := range encodedAuthors.Authors {
		authors = append(authors, map[string]any{
			"id":   author.Id,
			"name": author.Name,
		})
	}

	fmt.Printf("% x\n", authors)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(authors)
}

func (s *Server) AddAuthor(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
