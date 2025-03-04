package main

import (
	"consumer/src/api/controllers"
	"consumer/src/api/domain"
)

func main() {
	run()
}

func run() {
	srv := controllers.NewServer(domain.NewCore())
	srv.RabbitCh.MessageService.ReadMessageFromQueue()
}
