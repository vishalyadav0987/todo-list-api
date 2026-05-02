package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vishalyadav0987/todo-list-api/internal/config"
	"github.com/vishalyadav0987/todo-list-api/internal/infrastructure/db/sqlite"
)

func main() {
	fmt.Println("Working on Todo List Api")

	cfg := config.MustLoad()

	db := sqlite.NewConnection(cfg.DBPath)
	defer db.Close()

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	// 1. Testing Routes for Health Check
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Testing Routes for Health Check",
		})
	})

	// 2. Starting Server
	server := &http.Server{
		Addr: ":" + cfg.AppPort,
	}

	// 3. Run server in goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 4. Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")

}
