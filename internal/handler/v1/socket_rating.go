package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"work-project/internal/service"
)

func (h *Handler) initRatingSocket(router *gin.RouterGroup) {
	orders := router.Group("/all")
	{
		orders.GET("/", func(c *gin.Context) {
			h.SocketAggregator(c)
		})
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) SocketAggregator(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to establish WebSocket connection"})
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	var warehouseCodeUt string
	tokenChan := make(chan string)
	go func() {
		var msg struct {
			Token string `json:"token"`
		}
		if err := conn.ReadJSON(&msg); err != nil {
			service.GetHub().SendErrorToClient(&service.SocketClient{Conn: conn}, fmt.Errorf("failed to read token message"))
			close(tokenChan)
			return
		}
		tokenChan <- msg.Token
	}()

	var token string
	select {
	case token = <-tokenChan:
		if token == "" {
			service.GetHub().SendErrorToClient(&service.SocketClient{Conn: conn}, fmt.Errorf("the token was not received on time"))
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "missing token"))
			return
		}
	case <-time.After(5 * time.Second):
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "timeout waiting for token"))
		return
	}

	currentUser, err := h.GetUserFromToken(token)
	if err != nil {
		service.GetHub().SendErrorToClient(&service.SocketClient{Conn: conn}, fmt.Errorf("invalid token"))
		return
	}

	client := &service.SocketClient{
		Conn:     conn,
		UserID:   currentUser.UserID,
		UserName: currentUser.Email,
		Send:     make(chan service.SocketResponse, 256),
		Token:    token,
	}
	service.GetHub().Connect(client)
	log.Println("socket connected", "client", conn.RemoteAddr().String(), "count_of_connects", service.GetHub().GetCountOfClientsByGroup(warehouseCodeUt))

	h.serviceAggregator.SocketAggregate(context.Background(), conn, client, warehouseCodeUt)

}
