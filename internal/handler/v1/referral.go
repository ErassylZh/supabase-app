package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initReferral(v1 *gin.RouterGroup) {
	v1.GET(
		"/referral",
		middleware.GinErrorHandle(h.GetReferralCodeOfUser),
	)
	v1.POST(
		"/referral",
		middleware.GinErrorHandle(h.AddReferralCode),
	)
	v1.GET(
		"/referral/available",
		middleware.GinErrorHandle(h.GetAvailableReferralCodeOfUser),
	)
}

func (h *Handler) AddReferralCode(c *gin.Context) error {
	ctx := c.Request.Context()
	referralCode := c.Query("referralCode")
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	referral, err := h.usecases.Referral.AcceptReferralCode(ctx, userId, referralCode)
	if err != nil {
		return err
	}
	return schema.Respond(referral, c)
}

func (h *Handler) GetReferralCodeOfUser(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	referralCode, err := h.usecases.Referral.GetReferralCodeByUser(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(referralCode, c)
}

func (h *Handler) GetAvailableReferralCodeOfUser(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	referralCode, err := h.usecases.Referral.CheckAvailable(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(referralCode, c)
}
