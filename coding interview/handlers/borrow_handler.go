package handlers

import (
	"coding_interview/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BorrowHandler struct {
	BorrowService *services.BorrowService
}

func (h *BorrowHandler) BorrowBook(c *gin.Context) {
	var input struct {
		UserID uint `json:"user_id"`
		Days   int  `json:"days"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookID, _ := strconv.Atoi(c.Param("id"))
	err := h.BorrowService.BorrowBook(input.UserID, uint(bookID), input.Days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Book borrowed successfully"})
}
