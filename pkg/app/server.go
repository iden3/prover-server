package app

import (
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Server instance of chi server
type Server struct {
	Routes chi.Router
}

// NewServer creates new instance of server with routes
func NewServer(router chi.Router) *Server {
	return &Server{
		Routes: router,
	}
}

// Run starts the server
func (s *Server) Run(port int) {
	log.Infof("Server starting on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), s.Routes)
	log.Fatal(err)
}
