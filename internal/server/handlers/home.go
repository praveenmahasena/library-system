package handlers

import (
	"net/http"
)


func Home(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusOK)

}
