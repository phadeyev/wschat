package server

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

func New() *Server {
	router := chi.NewRouter()
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 2014,
	}
	serv := &Server{
		router:      router,
		upgrader:    upgrader,
		submutex:    &sync.Mutex{},
		subscribers: map[string]Subscriber{},
	}

	serv.bindRoutes()

	return serv
}

func (serv *Server) Start() error {
	return http.ListenAndServe(":8085", serv.router)
}
