package handlers

import (
	"context"
	"net/http"

	"github.com/praveenmahasena/libsys/internal/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	views.HomePage("")
}
