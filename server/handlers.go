package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (serv *Server) WShandler(w http.ResponseWriter, r *http.Request) {

	ws, err := serv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	go func() {
		for {
			<-time.After(5 * time.Second)
			msg := Message{
				Type: MTPing,
			}
			if err := ws.WriteJSON(msg); err != nil {
				log.Printf("ws send ping err: %v", err)
				break
			}
		}
	}()
	id := uuid.New().String()
	fmt.Println("client connected: ", id)
	serv.submutex.Lock()
	serv.subscribers[id] = func(msg string) error {
		m := Message{
			Type: MTMessage,
			Data: msg,
		}
		if err := ws.WriteJSON(m); err != nil {
			log.Printf("ws msg send err: %v", err)
		}
		return nil
	}
	serv.submutex.Unlock()

	for {
		msg := Message{}
		if err := ws.ReadJSON(&msg); err != nil {
			if !websocket.IsCloseError(err, 1001) {
				log.Println("ws msg read err: %v", err)
			}
			break
		}

		if msg.Type == MTPong {
			continue
		}

		if msg.Type == MTMessage {
			fmt.Println(msg.Data)
			serv.submutex.Lock()
			for _, sub := range serv.subscribers {
				if err := sub(msg.Data); err != nil {
					log.Println("ws msg subs err: %v", err)
				}
			}
			serv.submutex.Unlock()
		}
	}
	defer func() {
		ws.Close()
		serv.submutex.Lock()
		delete(serv.subscribers, id)
		serv.submutex.Unlock()
		fmt.Println("client disconnected: ", id)
	}()

}
