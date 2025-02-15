package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "backend_go/internal/models"
    "backend_go/internal/handlers"
)

func main() {
    // Initialize the database
    db, err := gorm.Open(sqlite.Open("words.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }

    // Migrate the schema
    models.InitializeDatabase(db)

    // Initialize Gin router
    r := gin.Default()

    // Define routes
    r.GET("/api/dashboard/last_study_session", handlers.GetLastStudySession)
    r.GET("/api/words", handlers.GetWords)
    r.GET("/api/groups", handlers.GetGroups)

    // Start the server
    r.Run(":8080")
}
