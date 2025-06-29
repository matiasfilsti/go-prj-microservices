package controllers

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router    *chi.Mux
	Templates map[string]*template.Template
}

func NewServer(router *chi.Mux) *Server {
	return &Server{
		Router: router,
	}
}

func (s *Server) LoadTemplates() error {
	pages := []string{"home.html", "login.html"}
	s.Templates = make(map[string]*template.Template)
	for _, page := range pages {
		tmpl, err := template.ParseFiles(
			filepath.Join(htmlTemplatesFolder, "base-layout.html"),
			filepath.Join(htmlTemplatesFolder, page),
		)
		if err != nil {
			return fmt.Errorf("error parsing %s: %w", page, err)
		}
		s.Templates[page] = tmpl
	}
	return nil
}
