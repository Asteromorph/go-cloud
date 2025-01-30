package server

import "canvas/handlers"

func (s *Server) setUpRoutes() {
	handlers.Health(s.mux)
}
