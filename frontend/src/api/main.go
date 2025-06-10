package main

import (
	"front/src/api/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	run()

}

func run() {
	router := chi.NewRouter()
	server := controllers.NewServer(router)
	err := server.LoadTemplates()
	if err != nil {
		panic("Error loading templates: " + err.Error())
	}
	server.ConfigureRouter()

	http.ListenAndServe(":8085", server.Router)
}
