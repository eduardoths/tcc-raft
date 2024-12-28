package handlers

import (
	"net/http"

	"github.com/eduardoths/tcc-raft/dto"
	"github.com/eduardoths/tcc-raft/pkg/client"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/gin-gonic/gin"
)

type DatabaseHandler struct {
	balancer *client.BalancerClient
	l        logger.Logger
}

func NewDatabaseHandler() *DatabaseHandler {
	dh := &DatabaseHandler{
		l: logger.MakeLogger("handler", "database", "server", "balancer"),
	}

	return dh
}

func (dh DatabaseHandler) Set(c *gin.Context) {
	var body dto.SetBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Error: dto.Point("Invalid body, should be JSON"),
		})
		return
	}

	resp, err := dh.balancer.Set(c.Request.Context(), body.ToArgs())
	if err != nil {
		c.JSON(http.StatusBadGateway, dto.Response{
			Error: dto.Point(err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Data: &resp,
	})
}

func (dh DatabaseHandler) Get(c *gin.Context) {
	param := c.Param("key")

	resp, err := dh.balancer.Get(c.Request.Context(), dto.GetArgs{
		Key: param,
	})

	if err != nil {
		c.JSON(http.StatusBadGateway, dto.Response{
			Error: dto.Point(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Data: resp.ToResponse(),
	})
}

func (dh DatabaseHandler) Delete(c *gin.Context) {
	params := c.Param("key")

	resp, err := dh.balancer.Delete(c.Request.Context(), dto.DeleteArgs{
		Key: params,
	})

	if err != nil {
		c.JSON(http.StatusBadGateway, dto.Response{
			Error: dto.Point(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Data: resp,
	})
}

func (dh *DatabaseHandler) GetLeader(c *gin.Context) {
	resp, err := dh.balancer.GetLeader(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadGateway, dto.Response{
			Error: dto.Point(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Data: resp,
	})
	return
}
