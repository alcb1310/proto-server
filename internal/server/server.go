package server

import (
	"fmt"
	"net/http"

	"github.com/alcb1310/proto-server/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	port uint
	DB   database.Database
}

func New(port uint) http.Server {
	srv := Server{
		port: port,
		DB:   database.New(),
	}

	return http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: srv.registerRoutes(),
	}
}

func (s Server) registerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", homeRoute)

	r.Route("/protobuf", func(r chi.Router) {
		r.Get("/authors", s.GetAllAuthors)
		r.Get("/books", s.GetAllBooks)
	})

	r.Route("/json", func(r chi.Router) {
		r.Get("/authors", s.JSONGetAllAuthors)
	})

	return r
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
