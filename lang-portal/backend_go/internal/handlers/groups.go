package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend_go/internal/service"
)

// GetGroups returns a list of groups
func GetGroups(c *gin.Context) {
    groups, err := service.GetGroups()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, groups)
}
