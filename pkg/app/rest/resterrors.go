package rest

import (
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
)

// ErrorJSON makes json and respond with error
func ErrorJSON(w http.ResponseWriter, r *http.Request, httpStatusCode int, err error, details string, errCode int) {
	logrus.Errorf("%d - %d - %v - %s", httpStatusCode, errCode, err, details)
	render.Status(r, httpStatusCode)
	render.JSON(w, r, map[string]interface{}{"code": errCode, "error": err.Error(), "details": details})
}
