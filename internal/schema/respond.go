package schema

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Result  T      `json:"result"`
}

func Respond[T any](v T, c *gin.Context) error {
	c.JSON(http.StatusOK, Response[T]{
		Status:  true,
		Message: "Success",
		Result:  v,
	})

	return nil
}

type Empty = struct{}

type Paginate[T any] struct {
	Items T     `json:"items"`
	Total int64 `json:"total"`
	Page  int   `json:"page"`
}

type ResponsePaginate[T any] struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     Paginate[T] `json:"result"`
}

func RespondPaginate[T any](v T, total int64, page int, c *gin.Context) error {
	c.JSON(http.StatusOK, ResponsePaginate[T]{
		Status:     true,
		Message:    "OK",
		StatusCode: 0,
		Result: Paginate[T]{
			Items: v,
			Total: total,
			Page:  page,
		},
	})
	return nil
}
