package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/dnogueir-org/seeker/internal/services"
	"github.com/dnogueir-org/seeker/web/handler"
	"github.com/gorilla/mux"
)

type Webserver struct {
	NavigationService services.NavigationService
}

func MakeNewWebServer(ns services.NavigationService) *Webserver {
	return &Webserver{
		NavigationService: ns,
	}
}

func (w Webserver) Serve() {
	router := mux.NewRouter()
	middleware := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeNavigationHandler(router, middleware, w.NavigationService)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
