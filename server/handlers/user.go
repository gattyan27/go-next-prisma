package handlers

import (
	"net/http"

	"github.com/gattyan27/go-next-prisma/prisma/client"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	prisma *client.PrismaClient
}

func NewUserHandler(prisma *client.PrismaClient) *UserHandler {
	return &UserHandler{prisma: prisma}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userId := c.Param("id")

	user, err := h.prisma.User.FindUnique(
		client.User.ID.Equals(userId),
	).Exec(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}