package handler

import (
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/dnogueir-org/seeker/internal/services"
	"github.com/dnogueir-org/seeker/web/common"
	"github.com/gorilla/mux"
)

func MakeNavigationHandler(router *mux.Router, n *negroni.Negroni, service services.NavigationService) {
	router.Handle("/navigation", n.With(
		negroni.Wrap(Navigate(service)),
	)).Methods("GET", "OPTIONS")
}

func Navigate(service services.NavigationService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "Application/json")

		sortingParam := r.URL.Query().Get("sorting")
		scoringProfileParam := r.URL.Query().Get("scoringProfile")
		fieldsParam := r.URL.Query()["field"]
		pageParam := r.URL.Query().Get("page")
		resultsPerPageParam := r.URL.Query().Get("resultsPerPage")

		page, err := strconv.Atoi(pageParam)
		if err != nil {
			common.BadRequest(w, "page param should be a number")
			return
		}

		resultsPerPage, err := strconv.Atoi(resultsPerPageParam)
		if err != nil {
			common.BadRequest(w, "resultsPerPage param should be a number")
			return
		}

		request := services.NavigationRequest{
			Sorting:        sortingParam,
			ScoringProfile: scoringProfileParam,
			Fields:         fieldsParam,
			Page:           page,
			ResultsPerPage: resultsPerPage,
		}

		navigationResponse, err := service.Navigate(request)
		if err != nil {
			common.InternalServerError(w, err.Error())
		}

		common.WriteJSONResponse(w, http.StatusOK, navigationResponse)
	})
}
