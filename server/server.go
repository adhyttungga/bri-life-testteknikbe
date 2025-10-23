package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adhyttungga/bri-life-testteknikbe/config"
	"github.com/adhyttungga/bri-life-testteknikbe/deliveries"
	"github.com/adhyttungga/bri-life-testteknikbe/middleware"
	"github.com/adhyttungga/bri-life-testteknikbe/repositories"
	"github.com/adhyttungga/bri-life-testteknikbe/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	mysqlDB *gorm.DB
}

func NewServer(mysqlDB *gorm.DB) *Server {
	return &Server{mysqlDB: mysqlDB}
}

func (s *Server) Run() error {
	// Initialize repositories
	agentRepo := repositories.NewAgentRepository(s.mysqlDB)
	transRepo := repositories.NewTransRepository(s.mysqlDB)
	productRepo := repositories.NewProductRepository(s.mysqlDB)
	// Initialize usecases
	authUC := usecases.NewAuthUsecase(agentRepo)
	agentUC := usecases.NewAgentUsecase(agentRepo)
	transUC := usecases.NewTransUsecase(transRepo, productRepo)
	// Initialize deliveries
	authDlvr := deliveries.NewAuthDelivery(authUC)
	agentDlvr := deliveries.NewAgentDelivery(agentUC)
	transDlvr := deliveries.NewTransDelivery(transUC)
	// Initialize gin server
	router := gin.Default()
	// Register routes
	authRoutes := router.Group("/api/v1/auth")
	{
		authRoutes.POST("/login", authDlvr.Login)
	}
	agentRoutes := router.Group("/api/v1/agents").Use(middleware.ProtectRoute())
	{
		agentRoutes.POST("/", agentDlvr.CreateAgent)
		agentRoutes.PUT("/", agentDlvr.UpdateAgent)
		agentRoutes.DELETE("/", agentDlvr.DeleteAgent)
	}
	transRoutes := router.Group("/api/v1/transaction").Use(middleware.ProtectRoute())
	{
		transRoutes.POST("/", transDlvr.CreateTrans)
		transRoutes.PUT("/", transDlvr.UpdateTrans)
		transRoutes.DELETE("/", transDlvr.DeleteTrans)
	}

	ctx, cancel := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    ":" + config.Config.Port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("failed to start server: %v", err)
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Printf("received signal: %v, shutting down server...", v)
	case done := <-ctx.Done():
		log.Printf("context done: %v, shutting down server...", done)
	}

	log.Println("Server exited properly")
	return nil
}
