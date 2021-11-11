package middleware

import (
	"github.com/go-chi/chi/middleware"
	"github.com/iden3/prover-server/pkg/log"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// ZapContextLogger is a middleware that logs the start and end of each request
func ZapContextLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		t1 := time.Now()
		defer func() {
			log.Info(r.Context(), "",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("remoteAddr", r.RemoteAddr),
				zap.Duration("lat", time.Since(t1)),
				zap.Int("status", ww.Status()))
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}
