package handler

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dnogueir-org/seeker/internal/services"
	"github.com/gorilla/mux"
)

type NavigationRequest struct{}

func MakeNavigationHandler(router *mux.Router, n *negroni.Negroni, service services.NavigationService) {
	router.Handle("/navigation", n.With(
		negroni.Wrap(Navigate(service)),
	)).Methods("GET", "OPTIONS")
}

func Navigate(service services.NavigationService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "Application/json")

		filterField := r.URL.Query()["field"]
		navigationResponse := service.Navigate(filterField)

		jsonResponse, err := responseBodyToJSON(navigationResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(jsonResponse)
	})
}
