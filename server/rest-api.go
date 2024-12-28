package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eduardoths/tcc-raft/handlers"
	"github.com/eduardoths/tcc-raft/internal/config"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RestApiServer struct {
	srv       *http.Server
	dbHandler *handlers.DatabaseHandler
	// adminHandler *handlers.AdminHandler
	log logger.Logger
}

func NewRestApiServer(log logger.Logger) (*RestApiServer, error) {
	cfg := config.Get()
	log = log.With(
		"server", "api",
		"port", cfg.Port,
	)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	api := router.Group("/api/v1")

	server := &RestApiServer{
		log:       log,
		dbHandler: handlers.NewDatabaseHandler(),
		// adminHandler: handlers.NewAdminHandler(),
	}
	log.Info("%d", cfg.Port)

	dbApi := api.Group("/db")
	dbApi.POST("/", server.dbHandler.Set)
	dbApi.GET("/:key", server.dbHandler.Get)
	dbApi.DELETE("/:key", server.dbHandler.Delete)
	dbApi.GET("/", server.dbHandler.GetLeader)

	// adminApi := api.Group("/admin")
	// adminApi.DELETE("/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	if err := server.adminHandler.ShutdownNode(id); err != nil {
	// 		c.JSON(http.StatusBadGateway, dto.Response{
	// 			Error: dto.Point(err.Error()),
	// 		})
	// 		return
	// 	}
	// 	server.dbHandler.RemoveNode(id)
	// 	c.JSON(http.StatusNoContent, dto.Response{})
	// })

	// adminApi.POST("/", func(c *gin.Context) {
	// 	var body dto.AddNodeArgs

	// 	if err := c.BindJSON(&body); err != nil {
	// 		c.JSON(http.StatusBadRequest, dto.Response{
	// 			Error: dto.Point(err.Error()),
	// 		})
	// 		return
	// 	}

	// 	if err := server.adminHandler.AddNode(body); err != nil {
	// 		c.JSON(http.StatusBadGateway, dto.Response{
	// 			Error: dto.Point(err.Error()),
	// 		})
	// 		return
	// 	}

	// 	if err := server.dbHandler.AddNode(body.ID, body.Addr()); err != nil {
	// 		log.Error(err, "failed to create client")
	// 	}

	// 	c.JSON(http.StatusCreated, dto.Response{
	// 		Data: "New node created",
	// 	})
	// })
	// adminApi.PATCH("/set-leader/", func(c *gin.Context) {
	// 	var body dto.SetLeader
	// 	if err := c.BindJSON(&body); err != nil {
	// 		c.JSON(http.StatusBadRequest, dto.Response{
	// 			Error: dto.Point(err.Error()),
	// 		})
	// 		return
	// 	}

	// 	if err := server.dbHandler.SetLeader(body); err != nil {
	// 		c.JSON(http.StatusBadGateway, dto.Response{
	// 			Error: dto.Point(err.Error()),
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusCreated, dto.Response{
	// 		Data: "set leader",
	// 	})
	// })

	server.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router.Handler(),
	}

	return server, nil
}

func (r *RestApiServer) Start() {
	go func() {
		if err := r.srv.ListenAndServe(); err != nil {
			r.log.Error(err, "failed to start server")
			panic(err)
		}
	}()
	defer r.Close()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func (r *RestApiServer) Close() {
	r.srv.Shutdown(context.Background())
	r.log.Info("Shuting down http server...")
}
