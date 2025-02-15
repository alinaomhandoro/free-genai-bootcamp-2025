package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend_go/internal/service"
)

// GetWords returns a list of words
func GetWords(c *gin.Context) {
    words, err := service.GetWords()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, words)
}
