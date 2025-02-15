// +build mage

package main

import (
    "fmt"
    "os"
    "os/exec"
)

// Initialize initializes the SQLite database
func Initialize() error {
    fmt.Println("Initializing database...")
    cmd := exec.Command("sqlite3", "words.db", ".databases")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

// Migrate runs the database migrations
func Migrate() error {
    fmt.Println("Running migrations...")
    cmd := exec.Command("sqlite3", "words.db", ".read db/migrations/0001_init.sql")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

// Seed seeds the database with initial data
func Seed() error {
    fmt.Println("Seeding database...")
    // Implement seeding logic here
    return nil
}
