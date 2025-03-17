package api

import (
	"UrlShortenerGoLang/services/url"
	"UrlShortenerGoLang/storage"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	urlHandler := url.NewHandler(s.store)
	urlHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, router)
}
