package handlers

import (
	"errors"
	"net/http"
	"sync"

	"github.com/eduardoths/tcc-raft/dto"
	"github.com/eduardoths/tcc-raft/internal/config"
	"github.com/eduardoths/tcc-raft/pkg/client"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/gin-gonic/gin"
)

type DatabaseHandler struct {
	mu      *sync.Mutex
	keys    []string
	index   int
	clients map[string]*client.DatabaseClient
	l       logger.Logger
}

func NewDatabaseHandler() *DatabaseHandler {
	dh := &DatabaseHandler{
		mu:      &sync.Mutex{},
		clients: make(map[string]*client.DatabaseClient),
		keys:    make([]string, 0),
		index:   -1,
		l:       logger.MakeLogger("handler", "database", "server", "balancer"),
	}

	servers := config.Get().RaftCluster.Servers

	for _, srv := range servers {
		client, err := client.NewDatabaseClient(srv.Addr())
		if err != nil {
			dh.l.Error(err, "Failed to generate new database client for server %s", srv.ID)
		}
		dh.clients[srv.ID] = client
		dh.keys = append(dh.keys, srv.ID)
	}

	return dh
}

func (dh *DatabaseHandler) next() (string, *client.DatabaseClient) {
	if len(dh.keys) == 0 {
		return "", nil
	}
	dh.index = (dh.index + 1) % len(dh.keys)
	key := dh.keys[dh.index]
	value := dh.clients[key]
	return key, value
}

func (dh *DatabaseHandler) AddNode(id string, addr string) error {
	client, err := client.NewDatabaseClient(addr)
	if err != nil {
		dh.l.Error(err, "Failed to generate new database client for server %s", id)
		return err
	}

	dh.mu.Lock()
	defer dh.mu.Unlock()
	dh.keys = append(dh.keys, id)
	dh.clients[id] = client
	return nil
}

func (dh *DatabaseHandler) RemoveNode(id string) error {
	dh.mu.Lock()
	defer dh.mu.Unlock()

	if _, exists := dh.clients[id]; exists {
		defer dh.clients[id].Close()
		delete(dh.clients, id)

		for i, key := range dh.keys {
			if id == key {
				dh.keys = append(dh.keys[:i], dh.keys[i+1:]...)
			}
		}

		if dh.index >= len(dh.keys) {
			dh.index = -1
		}
		return nil
	}

	return errors.New("server not found, cannot remove")
}

func (dh DatabaseHandler) Set(c *gin.Context) {
	var body dto.SetBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Error: dto.Point("Invalid body, should be JSON"),
		})
		return
	}

	id, client := dh.next()
	resp, err := client.Set(c.Request.Context(), body.ToArgs())
	if err != nil {
		dh.l.With("to-id", id).Error(err, "Failed to create set request")
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

	id, client := dh.next()

	resp, err := client.Get(c.Request.Context(), dto.GetArgs{
		Key: param,
	})

	if err != nil {
		dh.l.With("to-id", id).Error(err, "Failed to create get request")
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

	id, client := dh.next()

	err := client.Delete(c.Request.Context(), dto.DeleteArgs{
		Key: params,
	})

	if err != nil {
		dh.l.With("to-id", id).Error(err, "Failed to create delete request")
		c.JSON(http.StatusBadGateway, dto.Response{
			Error: dto.Point(err.Error()),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
