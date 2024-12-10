package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/eduardoths/tcc-raft/dto"
	"github.com/eduardoths/tcc-raft/handlers"
	"github.com/eduardoths/tcc-raft/internal/config"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	srv          *http.Server
	dbHandler    *handlers.DatabaseHandler
	adminHandler *handlers.AdminHandler
	log          logger.Logger
}

func NewHttpServer(log logger.Logger) (*HttpServer, error) {
	cfg := config.Get()
	log = log.With(
		"server", "api",
		"port", cfg.Port,
	)

	router := gin.Default()
	api := router.Group("/api/v1")

	server := &HttpServer{
		log:       log,
		dbHandler: handlers.NewDatabaseHandler(),
	}

	dbApi := api.Group("/db")
	dbApi.POST("/", server.dbHandler.Set)
	dbApi.GET("/:id", server.dbHandler.Get)
	dbApi.DELETE("/:id", server.dbHandler.Delete)

	adminApi := api.Group("/admin")
	adminApi.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := server.adminHandler.ShutdownNode(id); err != nil {
			c.JSON(http.StatusBadGateway, dto.Response{
				Error: dto.Point(err.Error()),
			})
			return
		}
		server.dbHandler.RemoveNode(id)
		c.JSON(http.StatusNoContent, dto.Response{})
	})

	adminApi.POST("/", func(c *gin.Context) {
		var body dto.AddNodeArgs

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, dto.Response{
				Error: dto.Point(err.Error()),
			})
			return
		}

		if err := server.adminHandler.AddNode(body); err != nil {
			c.JSON(http.StatusBadGateway, dto.Response{
				Error: dto.Point(err.Error()),
			})
			return
		}

		if err := server.dbHandler.AddNode(body.ID, body.Addr()); err != nil {
			log.Error(err, "failed to create client")
		}

		c.JSON(http.StatusCreated, dto.Response{
			Data: "New node created",
		})
	})

	server.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router.Handler(),
	}

	return server, nil
}

func (s *HttpServer) Start() {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			s.log.Error(err, "failed to start server")
			panic(err)
		}
	}()
	defer s.Close()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func (s *HttpServer) Close() {
	s.srv.Shutdown(context.Background())
	s.log.Info("Shuting down http server...")
}
