package main

import (
	"fmt"
	"ida_diag/src/database"
	"ida_diag/src/module/routes"
	"log"
	"net/http"
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
	db := database.InitDb(connStr)
	fmt.Println("Database connected")

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Addr: addr,
	}
	routeDefinitions := routes.RegisteredRoutes(db)
	for _, route := range routeDefinitions {
		http.HandleFunc(route.Path, route.Handle)
	}
	gin.SetMode(gin.ReleaseMode)

	go func() {
		fmt.Printf("Server is running on port %s\n", port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	select {}
}
