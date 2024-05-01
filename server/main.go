package main

import (
	"github.com/gattyan27/go-next-prisma/prisma/client"
	"github.com/gattyan27/go-next-prisma/server/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	prisma := client.NewClient()
	defer prisma.Disconnect()

	userHandler := handlers.NewUserHandler(prisma)

	router.POST("/user/:id", userHandler.GetUser)

	router.Run(":8080")
}