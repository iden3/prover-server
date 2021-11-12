package app

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/iden3/prover-server/pkg/log"
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
	log.Infow("Server started", "port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), s.Routes)
	log.Fatal(err)
}
