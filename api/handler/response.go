package handler

import (
	"encoding/json"
	"net/http"

	pkgerr "bitbucket.org/icehousecorp/moviedb/pkg/error"
)

func Write(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-type", "application/json")

	if e, ok := data.(*pkgerr.UnauthorizeError); ok {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(e)
	} else if e, ok := data.(*pkgerr.NotFoundError); ok {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(e)
	} else {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(data)
	}
}
