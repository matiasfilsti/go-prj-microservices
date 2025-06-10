package controllers

import (
	"net/http"
)

func (s *Server) PublicFiles(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/static/", http.FileServer(http.Dir("./controllers/html-templates/static"))).ServeHTTP(w, r)

}
