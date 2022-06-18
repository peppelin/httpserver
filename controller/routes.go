package controller

import "short2/handlers"

func (s *Server) routes() {
	s.router.HandleFunc("/", handlers.HandleIndex())
	s.router.HandleFunc("/admin-ok", handlers.HandleAdmin())
}
