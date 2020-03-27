package server

import (
	"sync"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

type MessageType string

const (
	MTPing    MessageType = "ping"
	MTPong    MessageType = "pong"
	MTMessage MessageType = "message"
)

type Message struct {
	Type MessageType `json:"type"`
	Data string      `json:"data,omitempty"`
}

type Subscriber func(msg string) error

type Server struct {
	router   *chi.Mux
	upgrader *websocket.Upgrader

	submutex    *sync.Mutex
	subscribers map[string]Subscriber
}
