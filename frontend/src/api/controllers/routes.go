package controllers

func (s *Server) ConfigureRouter() {

	s.Router.Get("/login", s.Login)
	s.Router.Get("/home", s.Home)
	s.Router.Get("/static/*", s.PublicFiles)

}
