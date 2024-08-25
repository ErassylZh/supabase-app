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
		"/marks",
		middleware.GinErrorHandle(h.CreateMark),
	)
	v1.GET(
		"/marks",
		middleware.GinErrorHandle(h.FindMarksByUserID),
	)
	v1.DELETE(
		"/marks",
		middleware.GinErrorHandle(h.DeleteMark),
	)
	v1.GET(
		"/marks",
		middleware.GinErrorHandle(h.FindPostsByUserID),
	)
}

func (h *Handler) CreateMark(c *gin.Context) error {
	ctx := c.Request.Context()
	var mark model.Mark
	if err := c.ShouldBindJSON(&mark); err != nil {
		return err
	}
	err := h.services.Mark.CreateMark(ctx, &mark)
	if err != nil {
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

func (h *Handler) FindPostsByUserID(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userID, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	posts, err := h.services.Mark.FindPostsByUserID(ctx, userID)
	if err != nil {
		return err
	}
	return schema.Respond(posts, c)
}
