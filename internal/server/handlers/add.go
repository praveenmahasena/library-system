package handlers

import (
	"context"
	"net/http"

	"github.com/praveenmahasena/libsys/internal/views"
)


func Add(w http.ResponseWriter,r *http.Request){
	views.Add().Render(context.TODO(),w)
	// error check is not done
}
