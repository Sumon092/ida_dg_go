


// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gin-gonic/gin"
// )

// func RegisterRoutes() {
// 	router := gin.Default()
// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"message": "Hello, world tumi koi!"})
// 	})
// 	router.GET("/profile", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"message": "This is the user profile"})
// 	})
// 	port := os.Getenv("PORT")
// 	addr := fmt.Sprintf(":%s", port)
// 	server := &http.Server{
// 		Addr:    addr,
// 		Handler: router,
// 	}

// 	go func() {
// 		fmt.Printf("Server is running on port %s\n", port)
// 		if err := server.ListenAndServe(); err != nil {
// 			log.Fatalf("Server error: %v", err)
// 		}
// 	}()
// }
