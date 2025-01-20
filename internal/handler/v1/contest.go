package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initContest(v1 *gin.RouterGroup) {
	v1.GET(
		"/contest",
		middleware.GinErrorHandle(h.GetAllHashtags),
	)
	v1.POST(
		"/join",
		middleware.GinErrorHandle(h.JoinContest),
	)
}

// GetActiveContest
// WhoAmi godoc
// @Summary получить активные контесты
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]schema.ContestData]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest [get]
func (h *Handler) GetActiveContest(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userID, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}

	contest, err := h.services.Contest.GetActive(ctx, userID)
	if err != nil {
		return err
	}
	return schema.Respond(contest, c)
}

// JoinContest
// WhoAmi godoc
// @Summary подключиться к розыгрышу
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[schema.Empty{}]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param data body schema.JoinContestRequest true "CreateMark"
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/join [post]
func (h *Handler) JoinContest(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userID, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}

	var data schema.JoinContestRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		return err
	}
	data.UserID = userID

	err = h.services.Contest.Join(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}
