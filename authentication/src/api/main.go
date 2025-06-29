package main

import (
	"authentication/src/api/controllers"
	"authentication/src/api/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	run()
}

func run() {
	core := domain.NewCore()
	srv := controllers.NewServer(gin.New(), core)
	srv.ConfigureRouter()
	srv.Router.Run(":8082")
}
