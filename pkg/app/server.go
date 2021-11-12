package app

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/iden3/prover-server/pkg/log"
	"go.uber.org/zap"
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
	log.Info("Server started", zap.Int("port", port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), s.Routes)
	log.Panic("server stopped", zap.Error(err))
}
