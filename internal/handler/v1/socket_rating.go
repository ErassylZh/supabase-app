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
		orders.GET("", func(c *gin.Context) {
			h.SocketAggregator(c)
		})
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SocketAggregator
// @Summary WebSocket for rating aggregation
// @Description Сокет для получения рейтинга всех участников в контесте (ws://host:port/ws/v1/all?)
// @Tags WebSocket
// @Accept json
// @Produce json
// @Param token body TokenData true "JWT Token"
// @Router /all [get]
func (h *Handler) SocketAggregator(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic recovered:", r)
		}
	}()

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to establish WebSocket connection"})
		return
	}
	defer conn.Close()

	// Таймаут для закрытия неактивных соединений
	pingTicker := time.NewTicker(30 * time.Second)
	defer pingTicker.Stop()

	// Таймер ожидания получения токена
	tokenChan := make(chan string)
	go func() {
		var msg TokenData
		conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // Таймаут на получение токена
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
	log.Println("socket connected", "client", conn.RemoteAddr().String(), "count_of_connects", service.GetHub().GetCountOfClientsByGroup("warehouseCodeUt"))

	// Контекст для управления завершением горутины
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Горутина для проверки активности соединения
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-pingTicker.C:
				err := conn.WriteMessage(websocket.PingMessage, nil)
				if err != nil {
					log.Println("closing inactive WebSocket connection:", conn.RemoteAddr().String())
					service.GetHub().Disconnect(client)
					conn.Close()
					return
				}
			}
		}
	}()

	h.serviceAggregator.SocketAggregate(ctx, conn, client, "")

}

type TokenData struct {
	Token string `json:"token"`
}
