package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend_go/internal/service"
)

// GetLastStudySession returns the most recent study session
func GetLastStudySession(c *gin.Context) {
    session, err := service.GetLastStudySession()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, session)
}
