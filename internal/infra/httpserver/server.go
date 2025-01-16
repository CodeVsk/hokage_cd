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
		Handlers: make(map[string]map[string]gin.HandlerFunc),
	}
}

func (h *HttpServer) RegisterHandler(route string, method string, handler gin.HandlerFunc) {
	if _, exists := h.Handlers[route]; !exists {
		h.Handlers[route] = make(map[string]gin.HandlerFunc)
	}
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