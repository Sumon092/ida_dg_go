package main

import (
	"fmt"
	"ida_diag/src/database"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	go func() {
		<-sigCh
		fmt.Println("\nServer is shutting down...")
		os.Exit(0)
	}()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connStr := os.Getenv("DATABASE_URL")
	database.InitDb(connStr)
	fmt.Println("Database connected")

	gin.SetMode(gin.ReleaseMode)
	RegisterRoutes()
	Chain()
	select {}
}
