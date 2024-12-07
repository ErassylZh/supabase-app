package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initMark(v1 *gin.RouterGroup) {
	v1.POST(
		"/mark",
		middleware.GinErrorHandle(h.CreateMark),
	)
	v1.GET(
		"/user/mark",
		middleware.GinErrorHandle(h.FindMarksByUserID),
	)
	v1.DELETE(
		"/mark/:post_id",
		middleware.GinErrorHandle(h.DeleteMark),
	)
}

// CreateMark
// WhoAmi godoc
// @Summary сохранить в избранное
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.Mark]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param data body schema.CreateMark true "CreateMark"
// @Security BearerAuth
// @tags mark
// @Router /api/v1/mark [post]
func (h *Handler) CreateMark(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userID, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	var mark schema.CreateMark
	if err := c.ShouldBindJSON(&mark); err != nil {
		return err
	}
	mark.UserID = userID
	if err := h.services.Mark.CreateMark(ctx, mark); err != nil {
		return err
	}
	return schema.Respond(mark, c)
}

// FindMarksByUserID
// WhoAmi godoc
// @Summary все избранное пользователя
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]schema.PostResponse]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @Param filter query string true "all, post, partner"
// @tags mark
// @Router /api/v1/user/mark [get]
func (h *Handler) FindMarksByUserID(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userID, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	filter := c.Query("filter")

	marks, err := h.services.Mark.FindPostsByUserID(ctx, userID, filter)
	if err != nil {
		return err
	}
	return schema.Respond(marks, c)
}

// DeleteMark
// WhoAmi godoc
// @Summary удалить в избранное
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.Mark]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags mark
// @Router /api/v1/mark/:post_id [delete]
func (h *Handler) DeleteMark(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userID, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	postIdsStr := c.Param("post_id")

	postId, err := strconv.ParseUint(postIdsStr, 10, 64)
	if err != nil {
		return err
	}
	if err := h.services.Mark.DeleteMark(ctx, uint(postId), userID); err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}
