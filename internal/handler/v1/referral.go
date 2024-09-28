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

// AddReferralCode
// WhoAmi godoc
// @Summary пометить прочитанным сторис
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.ReferralCode]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param referralCode query string true "referralCode"
// @Security BearerAuth
// @tags referral
// @Router /api/v1/referral [post]
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

// GetReferralCodeOfUser
// WhoAmi godoc
// @Summary получить реферал код юзера
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.ReferralCode]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags referral
// @Router /api/v1/referral [get]
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

// GetAvailableReferralCodeOfUser
// WhoAmi godoc
// @Summary получить активность рефералки
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[schema.CheckAvailable]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags referral
// @Router /api/v1/referral/available [get]
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
