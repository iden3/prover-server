package middleware

import (
	"github.com/go-chi/chi/middleware"
	"github.com/iden3/prover-server/pkg/log"
	"net/http"
	"time"
)

// ZapContextLogger is a middleware that logs the start and end of each request
func ZapContextLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		t1 := time.Now()
		defer func() {
			log.WithContext(r.Context()).Infow("http	",
				"method", r.Method,
				"path", r.URL.Path,
				"remoteAddr", r.RemoteAddr,
				"lat", time.Since(t1),
				"status", ww.Status())
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}
