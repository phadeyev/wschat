package server

import (
	"net/http"
)

func (serv *Server) bindRoutes() {
	serv.router.Handle("/*", http.FileServer(http.Dir("./www")))
	serv.router.Get("/ws", serv.WShandler)
}
