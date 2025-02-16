package middlewares

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware логирует запросы и ответы сервера
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/swagger/" || len(c.Request.URL.Path) >= 9 && c.Request.URL.Path[:9] == "/swagger/" {
			c.Next()
			return
		}

		startTime := time.Now()

		// Логируем запрос

		if c.Request.Method == http.MethodGet {
			log.Printf("➡️ REQUEST: %s %s\nHeaders: %v\n",
				c.Request.Method, c.Request.URL, c.Request.Header)
		} else {
			body, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

			log.Printf("➡️ REQUEST: %s %s\nHeaders: %v\nBody: %s\n",
				c.Request.Method, c.Request.URL, c.Request.Header, string(body))

			responseBody := &bytes.Buffer{}
			writer := &responseRecorder{ResponseWriter: c.Writer, body: responseBody}
			c.Writer = writer

			// Выполняем запрос
			c.Next()

			// Логируем ответ
			log.Printf("⬅️ RESPONSE: %d\nBody: %s\nTime: %v\n",
				c.Writer.Status(), responseBody.String(), time.Since(startTime))
		}
	}
}

// responseRecorder оборачивает ResponseWriter, чтобы перехватывать тело ответа
type responseRecorder struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b) // Записываем в буфер
	return r.ResponseWriter.Write(b)
}
