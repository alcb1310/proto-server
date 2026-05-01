package server

import (
	"fmt"
	"net/http"

	"github.com/alcb1310/proto-server/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", homeRoute)

	r.Route("/protobuf", func(r chi.Router) {
		r.Route("/authors", func(r chi.Router) {
			r.Post("/", s.AddAuthor)
			r.Get("/", s.GetAllAuthors)
		})

		r.Get("/books", s.GetAllBooks)
	})

	r.Route("/json", func(r chi.Router) {
		r.Get("/authors", s.JSONGetAllAuthors)
		r.Get("/books", s.JSONGetAllBooks)
	})

	return r
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
