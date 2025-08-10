package main

import (
	"aioperator/src/api/controllers"
	"aioperator/src/api/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	run()
}

func run() {
	srv := controllers.NewServer(gin.New(), domain.NewCore())
	srv.ConfigureRouter()
	srv.Router.Run(":8086")
}
