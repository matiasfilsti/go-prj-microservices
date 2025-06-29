package controllers

import (
	"front/src/api/config"
	"net/http"
)

var htmlTemplatesFolder = "html-templates/"

var data map[string]interface{} = map[string]interface{}{
	"BackendUrl":  config.BackEndUrl,
	"BackendPort": config.BackEndPort,
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	tmpl, ok := s.Templates["login.html"]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, ok := s.Templates["home.html"]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
