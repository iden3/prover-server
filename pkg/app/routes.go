package app

import (
	"github.com/iden3/prover-server/pkg/app/handlers"
	customMiddleware "github.com/iden3/prover-server/pkg/app/middleware"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// Handlers contain supported handlers by server
type Handlers struct {
	/* Put handlers here*/
	ZKHandler *handlers.ZKHandler
}

// Routes initializes router
func (s *Handlers) Routes() chi.Router {

	r := chi.NewRouter()

	// Basic CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})
	r.Use(corsHandler.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(customMiddleware.ZapContextLogger)
	r.Use(middleware.Recoverer)
	r.Route("/api/v1", func(api chi.Router) {

		api.Use(render.SetContentType(render.ContentTypeJSON))

		api.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, struct {
				Status string `json:"status"`
			}{Status: "up and running"})
		})

		// identity routes, require auth and admin users only
		api.Route("/proof", func(rr chi.Router) {
			rr.Post("/generate", s.ZKHandler.GenerateProof)
			rr.Post("/verify", s.ZKHandler.VerifyProof)
		})

	})

	return r
}
