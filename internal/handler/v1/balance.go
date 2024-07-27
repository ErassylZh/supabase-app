package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initBalance(v1 *gin.RouterGroup) {
	v1.GET(
		"/balance",
		middleware.GinErrorHandle(h.GetBalanceOfUser),
	)
	v1.GET(
		"/balance/history",
		middleware.GinErrorHandle(h.GetHistoryOfTransactions),
	)
}

func (h *Handler) GetBalanceOfUser(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	balance, err := h.services.Balance.GetByUserId(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(balance, c)
}

func (h *Handler) GetHistoryOfTransactions(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	balance, err := h.services.Balance.GetTransactionHistory(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(balance, c)
}
