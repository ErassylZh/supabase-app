package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrResponse struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}

func (e ErrResponse) ToJson() (raw []byte, err error) {
	raw, err = json.Marshal(e)
	return raw, err
}
func GinErrorHandle(h func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			if len(c.Errors) == 0 {
				c.Error(err)
			}

			GinRecoveryFn(c)
		}
	}
}

func GinRecoveryFn(c *gin.Context) {
	err := c.Errors.Last()
	if err == nil {
		return
	}

	resp := &ErrResponse{
		Status:     false,
		StatusCode: 999999,
		Message:    err.Err.Error(),
		Result:     struct{}{},
	}
	if c.IsAborted() {
		rawResp, _ := resp.ToJson()
		c.Writer.Write(rawResp)
		return
	}

	if err.Type == gin.ErrorTypeBind {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusInternalServerError, resp)
}
