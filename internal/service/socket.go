package service

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

var hubGlobal *Hub

// Делаем такой костыль, потому что фронт привязывается к status_code и не может распарсить объект как массив или наоборот
type StatusCodeType int

const SocketResponseArrayStatusCode StatusCodeType = 0
const SocketResponseObjectStatusCode StatusCodeType = 1

type SocketResponse struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}

type Hub struct {
	Groups     map[string]*ClientGroup
	Register   chan *SocketClient
	Unregister chan *SocketClient
}

type SocketClient struct {
	GroupID  string
	UserID   string
	UserName string
	Conn     *websocket.Conn
	Send     chan SocketResponse
	Token    string
}

type ClientGroup struct {
	GroupID   string
	Clients   map[*SocketClient]bool
	Broadcast chan interface{}
}

func NewHub() {
	hubGlobal = &Hub{
		Groups:     make(map[string]*ClientGroup),
		Register:   make(chan *SocketClient),
		Unregister: make(chan *SocketClient),
	}
	go hubGlobal.Run()
}

func GetHub() *Hub {
	return hubGlobal
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Unregister:
			h.RemoveClientFromGroup(client, client.GroupID)
		case client := <-h.Register:
			h.AddClientToGroup(client, client.GroupID)
		}
	}
}

func (h *Hub) AddClientToGroup(client *SocketClient, groupID string) {
	if _, ok := h.Groups[groupID]; !ok {
		group := &ClientGroup{
			GroupID:   groupID,
			Clients:   make(map[*SocketClient]bool),
			Broadcast: make(chan interface{}),
		}
		h.Groups[groupID] = group
		go group.Run()
	}
	h.Groups[groupID].Clients[client] = true
}

func (h *Hub) RemoveClientFromGroup(client *SocketClient, groupID string) {
	if group, ok := h.Groups[groupID]; ok {
		if _, exists := group.Clients[client]; exists {
			delete(group.Clients, client)
			if len(group.Clients) == 0 {
				delete(h.Groups, groupID)
			}
		}
	}
}

func (h *Hub) SendToGroup(groupID string, message interface{}) {
	if group, ok := h.Groups[groupID]; ok {
		group.Broadcast <- message
	}
}

func (h *Hub) GetConnectedUsersByGroup(groupID string) []SocketClient {
	var result []SocketClient
	if len(h.Groups) == 0 {
		return nil
	}
	for client, _ := range h.Groups[groupID].Clients {
		result = append(result, *client)
	}
	return result
}

func (h *Hub) Connect(client *SocketClient) {
	GetHub().Register <- client
}

func (h *Hub) GetCountOfClientsByGroup(groupId string) int {
	group, exist := h.Groups[groupId]
	if !exist {
		return 0
	}
	return len(group.Clients)
}

func (h *Hub) SendDelayedToClient(client *SocketClient, message interface{}, delay time.Duration) {
	time.AfterFunc(delay, func() {
		client.Send <- SocketResponse{
			Status:     true,
			StatusCode: int(SocketResponseArrayStatusCode),
			Message:    "Delayed notification received",
			Result:     message,
		}
	})
}

func (h *Hub) SendMessageToClient(client *SocketClient, statusCode StatusCodeType, message interface{}) {
	select {
	case client.Send <- SocketResponse{
		Status:     true,
		StatusCode: int(statusCode),
		Message:    "Notification received",
		Result:     message,
	}:
	default:
		log.Println("Failed to send message to client", "client", client.UserID, "message", message)
	}
}

func (h *Hub) SendErrorToClient(client *SocketClient, err error) {
	select {
	case client.Send <- SocketResponse{
		Status:  false,
		Message: "Notification received with error",
		Result:  err.Error(),
	}:
	default:
		log.Println("Failed to send error message to client", "client", client.UserID, "err", err)
	}
}

func (cg *ClientGroup) Run() {
	for {
		select {
		case message := <-cg.Broadcast:
			for client := range cg.Clients {
				select {
				case client.Send <- SocketResponse{
					Status:  true,
					Message: "Notification received",
					Result:  message,
				}:
				default:
					close(client.Send)
					delete(cg.Clients, client)
				}
			}
		}
	}
}
