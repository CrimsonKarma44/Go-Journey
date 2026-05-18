package handlers

import (
	"coding_interview/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookHandler struct {
	BookService *services.BookService
}

func (h *BookHandler) ListBooks(c *gin.Context) {
	books, err := h.BookService.ListBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.BookService.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) FilterBooks(c *gin.Context) {
	publisher := c.Query("publisher")
	category := c.Query("category")
	books, err := h.BookService.FilterBooks(publisher, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}
