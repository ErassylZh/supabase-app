package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"work-project/internal/middleware"
	"work-project/internal/model"
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
		"/mark/:mark_id",
		middleware.GinErrorHandle(h.DeleteMark),
	)
}

func (h *Handler) CreateMark(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userID, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	var mark model.Mark
	if err := c.ShouldBindJSON(&mark); err != nil {
		return err
	}
	mark.UserID = userID
	if err := h.services.Mark.CreateMark(ctx, mark); err != nil {
		return err
	}
	return schema.Respond(mark, c)
}

func (h *Handler) FindMarksByUserID(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userID, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	marks, err := h.services.Mark.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}
	return schema.Respond(marks, c)
}

func (h *Handler) DeleteMark(c *gin.Context) error {
	ctx := c.Request.Context()
	markID, err := strconv.ParseUint(c.Param("mark_id"), 10, 64)
	if err != nil {
		return err
	}
	if err := h.services.Mark.DeleteMark(ctx, uint(markID)); err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}
