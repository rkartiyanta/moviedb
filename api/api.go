package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"bitbucket.org/icehousecorp/moviedb/api/handler/movie"
	"bitbucket.org/icehousecorp/moviedb/core"

	"github.com/go-chi/chi"
)

func New(
	movieStore core.MovieStore,
) *Server {
	return &Server{
		movieStore: movieStore,
	}
}

type Server struct {
	movieStore core.MovieStore
}

func (s *Server) handler() http.Handler {
	r := chi.NewRouter()
	r.Route("/movie", func(r chi.Router) {
		r.Get("/", movie.Request(s.movieStore))
	})
	return r
}

func (s *Server) Run() {
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", 8089),
		Handler: s.handler(),
	}
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("server shutdown: %v", err)
		}
	}()
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("could not start server: %v", err)
	}
}
