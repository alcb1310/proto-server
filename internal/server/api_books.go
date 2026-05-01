package server

import (
	"encoding/json"
	"fmt"
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

	fmt.Printf("% x\n", resp)

	w.Header().Set("Content-Type", "application/protobuf")
	w.Write(resp)
}

func (s *Server) JSONGetAllBooks(w http.ResponseWriter, r *http.Request) {
	encodedBooks, err := s.DB.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	books := make([]map[string]any, 0)
	for i, book := range encodedBooks.Books {
		books = append(books, map[string]any{
			"id":        book.Id,
			"title":     book.Title,
			"publisher": book.Publisher,
			"genre":     book.Genre,
		})

		books[i]["authors"] = make([]map[string]any, 0)

		for _, author := range book.Authors {
			currentAuthor := map[string]any{
				"id":   author.Id,
				"name": author.Name,
			}

			books[i]["authors"] = append(books[i]["authors"].([]map[string]any), currentAuthor)
		}
	}

	fmt.Printf("% x\n", books)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(books)
}
