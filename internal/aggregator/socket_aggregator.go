package aggregator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"work-project/internal/schema"
	"work-project/internal/service"
)

type SocketTopic string

const (
	SOCKET_TOPIC_ERROR  SocketTopic = "error"
	SOCKET_TOPIC_RATING SocketTopic = "rating"
)

type ServiceAggregator interface {
	SocketAggregate(ctx context.Context, c *websocket.Conn, client *service.SocketClient, warehouseCodeUt string)
}

type ServiceAggregatorService struct {
	services service.Services
}

func NewServiceAggregatorService(service service.Services) *ServiceAggregatorService {
	return &ServiceAggregatorService{
		services: service,
	}
}

func (s *ServiceAggregatorService) SocketAggregate(ctx context.Context, c *websocket.Conn, client *service.SocketClient, warehouseCodeUt string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic recovered:", r)
		}
	}()

	ticker := time.NewTicker(10 * time.Second)
	data, err := s.getActiveMessages(ctx, client)
	if err != nil {
		return
	}
	service.GetHub().SendMessageToClient(client, service.SocketResponseArrayStatusCode, data)

	go func(c *websocket.Conn, client *service.SocketClient) {
		defer func() {
			service.GetHub().Unregister <- client
			close(client.Send)
			c.Close()
		}()

		for {
			_, messageRaw, err := c.ReadMessage()
			if err != nil {
				return
			}
			s.handleClientMessage(ctx, client, messageRaw)
		}
	}(c, client)

	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				_ = c.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(10*time.Second))
				log.Println("client send channel closed")
				return
			}

			_ = c.SetWriteDeadline(time.Now().Add(90 * time.Second))

			if err := c.WriteJSON(message); err != nil {
				log.Printf("error while write message %s", err.Error())
				return
			}
		case <-ticker.C:
			_ = c.SetWriteDeadline(time.Now().Add(15 * time.Second))
			if err := c.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second)); err != nil {
				service.GetHub().SendErrorToClient(client, fmt.Errorf("error while ping connect"))
				log.Fatal("error while ping connect", "err", err)
				return
			}

			data, err = s.getActiveMessages(ctx, client)
			if err != nil {
				log.Printf("error get active messages %s", err.Error())
				return
			}
			service.GetHub().SendMessageToClient(client, service.SocketResponseArrayStatusCode, data)
		}
	}
}

func (s *ServiceAggregatorService) handleClientMessage(ctx context.Context, client *service.SocketClient, message []byte) {
	var msg schema.SocketMessage
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Fatal("invalid message format", "error", err)
		return
	}
	switch msg.SocketTopic {
	case string(SOCKET_TOPIC_RATING):
		data, err := s.getActiveMessages(ctx, client)
		if err != nil {
			log.Fatal(ctx, "error get active messages", "err", err)
			return
		}
		service.GetHub().SendMessageToClient(client, service.SocketResponseArrayStatusCode, data)
	}

}

func (s *ServiceAggregatorService) getActiveMessages(ctx context.Context, client *service.SocketClient) (interface{}, error) {
	result := make([]interface{}, 0)

	contestData, err := s.services.Contest.GetDataForSocket(ctx, client.UserID)
	if err != nil {
		return nil, err
	}

	for _, contest := range contestData {
		result = append(result, contest)
	}

	//sort.Slice(result, func(i, j int) bool {
	//	t1 := result[i].(schema.SocketResponse).Data.(schema.HasCreatedAt).GetCreatedAt()
	//	t2 := result[j].(schema.SocketResponse).Data.(schema.HasCreatedAt).GetCreatedAt()
	//	return t1.Before(t2)
	//})
	return result, nil
}

func (s *ServiceAggregatorService) appendToResult(result *[]interface{}, data interface{}, socketTopic SocketTopic) {
	*result = append(*result, schema.SocketResponse{
		Data:        data,
		SocketTopic: string(socketTopic),
	})
}
