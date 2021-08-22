package web

import (
	"eud_backup/pkg/web/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

// Server represents the HTTP server
type Server struct {
	*http.Server
	handlers []Handler
	router   *mux.Router
}

// Handler represents the HTTP function handler
type Handler struct {
	Path     string
	Methods  []string
	Function http.HandlerFunc
}

// Start starts the HTTP server
func (server Server) Start() error {
	server.registerRoutes()
	return server.ListenAndServe()
}

// registerRoutes registers all registered server handlers
func (server *Server) registerRoutes() {
	server.loadRoutes()
	for _, route := range server.handlers {
		server.router.HandleFunc(route.Path, route.Function).Methods(route.Methods...)
	}
	server.updateServerHandler()
}

// updateServerHandler updates the default HTTP handler to the custom Mux router which will handle all the connections
func (server *Server) updateServerHandler() {
	server.Handler = server.router
}

// loadRoutes loads all handlers
func (server *Server) loadRoutes() {
	server.handlers = []Handler{
		{
			"/",
			[]string{"GET"},
			handlers.Stats,
		},
	}
}


