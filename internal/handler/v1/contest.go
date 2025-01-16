package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
)

func (h *Handler) initContest(v1 *gin.RouterGroup) {
	v1.GET(
		"/contest",
		middleware.GinErrorHandle(h.GetAllHashtags),
	)
}

// GetActiveContest
// WhoAmi godoc
// @Summary получить все коллекций
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.Hashtag]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest [get]
func (h *Handler) GetActiveContest(c *gin.Context) error {
	//ctx := c.Request.Context()
	//token := c.GetHeader("Authorization")
	//userID, err := h.services.Auth.VerifyToken(token)
	//if err != nil {
	//	return err
	//}
	//
	//contest, err := h.services.Hashtag.GetAll(ctx)
	//if err != nil {
	//	return err
	//}
	//return schema.Respond(contest, c)
	return nil
}
