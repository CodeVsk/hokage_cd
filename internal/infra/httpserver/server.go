package httpserver

import (
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Server *gin.Engine
	Addr string
	Handlers map[string]map[string]gin.HandlerFunc 
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		Server: gin.Default(),
		Addr: addr,
	}
}

func (h *HttpServer) RegisterHandler(method string, route string, handler gin.HandlerFunc) {
	h.Handlers[route][method] = handler
}

func (h *HttpServer) Start() {
	for route, methods := range h.Handlers {
		for method, handler := range methods {
			h.Server.Handle(method, route, handler)
		}
	}

	h.Server.Run(h.Addr)
}