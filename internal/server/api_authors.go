package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pb "github.com/alcb1310/proto-server/cmd/schemas"
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	author := &pb.Author{}
	if err := proto.Unmarshal(body, author); err != nil {
		http.Error(w, fmt.Sprintf("failed to unmarshal body: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("%s\n", author.Name)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Created author with name: %s", author.Name)
}
