package schema

import "time"

type SocketResponse struct {
	SocketTopic string      `json:"socket_topic"`
	Data        interface{} `json:"data"`
}

type SocketMessage struct {
	SocketTopic string      `json:"socket_topic"`
	Message     interface{} `json:"message"`
}

type SocketErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type HasCreatedAt interface {
	GetCreatedAt() time.Time
}
