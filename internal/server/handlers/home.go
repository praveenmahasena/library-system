package handlers

import (
	"context"
	"net/http"

	"github.com/praveenmahasena/libsys/internal/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	comp := views.TemplateName("yayaya")
	comp.Render(context.Background(), w)
}
