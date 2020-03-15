package main

func (s *Server) routes() {
	s.router.Get("/users", chain(s.logger, s.respond, s.getUsers))
}