package api

import (
	"net/http"

	"github.com/gorilla/mux"

	v1 "financial-app-backend/internal/api/v1"
	"financial-app-backend/internal/database"
)

//NewRouter provide a handler API service.
func NewRouter(db database.Database) (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)

	return router, nil
}
