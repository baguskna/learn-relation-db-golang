package main

import (
	"learn-relation-db-go/config"
	"learn-relation-db-go/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/order", handler.CreateOrder)
	r.POST("/user", handler.CreateUser)
	r.GET("/user-with-order", handler.GetUserWithOrders)
	config.StartDB()
	r.Run()
}
