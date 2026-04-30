package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	port uint
}

func New(port uint) http.Server {
	srv := Server{
		port: port,
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

	return r
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
