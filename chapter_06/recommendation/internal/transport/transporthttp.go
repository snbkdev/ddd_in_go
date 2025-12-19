package transport

import (
	"DDD/chapter_06/recommendation/internal/recommendation"
	"net/http"

	"github.com/gorilla/mux"
)

func NewMux(recHandler recommendation.Handler) *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/recommendation", recHandler.GetRecommendation).Methods(http.MethodGet)
	return m
}