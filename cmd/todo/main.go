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
	httpapp "github.com/vishalyadav0987/todo-list-api/interfaces/http"
	"github.com/vishalyadav0987/todo-list-api/interfaces/http/handler"
	authapp "github.com/vishalyadav0987/todo-list-api/internal/application/auth"
	"github.com/vishalyadav0987/todo-list-api/internal/config"
	"github.com/vishalyadav0987/todo-list-api/internal/infrastructure/hasher"
	"github.com/vishalyadav0987/todo-list-api/internal/infrastructure/id"
	"github.com/vishalyadav0987/todo-list-api/internal/infrastructure/persistence/sqlite"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Working on Todo List Api")

	cfg := config.MustLoad()

	db := sqlite.NewConnection(cfg.DBPath)
	defer db.Close()

	// Skip migration till now
	// 0.8 migration
	// sqlite.RunMigration(db, cfg.DBPath)

	// router := gin.New()
	userRepo := sqlite.NewUserRepository(db)
	hasher := hasher.NewBcryptHasher(bcrypt.DefaultCost)
	idGen := id.NewUUIDGenerator()
	registerUc := authapp.NewRegisterUsecase(userRepo, hasher, idGen)

	authHandler := handler.NewAuthHnadler(registerUc)
	router := httpapp.SetUpRouter(authHandler)

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
		Addr:    ":" + cfg.AppPort,
		Handler: router,
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
