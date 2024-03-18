package http

import (
	"net/http"
)

type Server struct {
	httpServer http.Server
}

func NewServer(host string) Server {
	return Server{
		httpServer: http.Server{
			Addr: host,
		},
	}
}

func (s *Server) Run(mux *http.ServeMux) error {
	return http.ListenAndServe(s.httpServer.Addr, mux)
}
