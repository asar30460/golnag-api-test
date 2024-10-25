package main

import (
	"example.com/golang-api-test/server"

	"github.com/gin-gonic/gin"
)

func main() {
	server := server.NewGeneralServer()

	go server.BroadcastMessage()

	r := gin.Default()
	r.GET("/ws", func(ctx *gin.Context) {
		server.HandleWS(ctx)
	})
	r.Run(":3000")
}
