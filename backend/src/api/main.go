package main

import (
	"backend/src/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	run()
}

func run() {
	srv := controllers.NewServer(gin.New())
	srv.ConfigureRouter()
	srv.Router.Run(":8081")
}
