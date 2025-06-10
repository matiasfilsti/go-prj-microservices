package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

var htmlTemplatesFolder = "controllers/html-templates/"
var pageTemplates = map[string]*template.Template{}

// var htmlLayoutFolder = "controllers/html-layout/"

// var templ = template.Must(template.ParseFiles("controllers/html-templates/login.html", "controllers/html-templates/home.html", "controllers/html-templates/base-layout.html"))

// func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("accediendo")

// 	templ, err := template.ParseFiles("controllers/html-templates/base-layout.html", "controllers/html-templates/home.html")
// 	fmt.Println(err)
// 	fmt.Println(templ)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	templ.Execute(w, nil)
// }

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	tmpl, ok := s.Templates["login.html"]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	err := tmpl.ExecuteTemplate(w, "base", nil)
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

func (s *Server) htmlRendered(htmlFile string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(htmlTemplatesFolder+htmlFile, htmlTemplatesFolder+"base-layout.html")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return tmpl, nil
}

func (s *Server) Home2(w http.ResponseWriter, r *http.Request) {
	tmpl, ok := pageTemplates["home.html"]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// func LoadTemplates() error {
// 	pages := []string{"home.html", "login.html"} // agrega aquí más páginas si tienes
// 	for _, page := range pages {
// 		tmpl, err := template.ParseFiles(
// 			filepath.Join(htmlTemplatesFolder, "base-layout.html"),
// 			filepath.Join(htmlTemplatesFolder, page),
// 		)
// 		if err != nil {
// 			return fmt.Errorf("error parsing %s: %w", page, err)
// 		}
// 		pageTemplates[page] = tmpl
// 	}
// 	return nil
// }
