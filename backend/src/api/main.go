package main

import (
	"backend/src/api/client"
	"backend/src/api/controllers"
	"backend/src/api/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	run()
}

func run() {
	httpclient := client.CreateHttpClient()
	hclientController := client.NewClientHttp(httpclient)
	srv := controllers.NewServer(gin.New(), domain.NewCore(), hclientController)
	srv.ConfigureRouter()
	srv.Router.Run(":8081")
}
