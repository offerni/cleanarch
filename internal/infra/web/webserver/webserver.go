package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        *chi.Mux
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	if _, exists := s.Handlers[path]; !exists {
		s.Handlers[path] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[path][method] = handler
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, methods := range s.Handlers {
		for method, handler := range methods {
			switch method {
			case http.MethodGet:
				s.Router.Get(path, handler)
			case http.MethodPost:
				s.Router.Post(path, handler)
			default:
				s.Router.Handle(path, handler)
			}
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
