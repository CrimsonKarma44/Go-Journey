package handlers

import (
	"coding_interview/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}

func (h *UserHandler) EnrollUser(c *gin.Context) {
	var input struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.EnrollUser(input.Email, input.FirstName, input.LastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
