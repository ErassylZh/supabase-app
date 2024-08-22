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
		middleware.GinErrorHandle(h.Create),
	)
	v1.GET(
		"/marks",
		middleware.GinErrorHandle(h.FindMarksByUserID),
	)
	v1.DELETE(
		"/marks",
		middleware.GinErrorHandle(h.Delete),
	)
	v1.GET(
		"/marks",
		middleware.GinErrorHandle(h.FindPostsByUserID),
	)
}

func (h *Handler) Create(c *gin.Context) error {
	ctx := c.Request.Context()
	var mark model.Mark
	if err := c.ShouldBindJSON(&mark); err != nil {
		return err
	}
	err := h.services.Mark.Create(ctx, &mark)
	if err != nil {
		return err
	}
	return schema.Respond(mark, c)
}

func (h *Handler) FindMarksByUserID(c *gin.Context) error {
	ctx := c.Request.Context()
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return err
	}
	marks, err := h.services.Mark.FindByUserID(ctx, uint(userID))
	if err != nil {
		return err
	}
	return schema.Respond(marks, c)
}

func (h *Handler) Delete(c *gin.Context) error {
	ctx := c.Request.Context()
	markID, err := strconv.ParseUint(c.Param("mark_id"), 10, 64)
	if err != nil {
		return err
	}
	if err := h.services.Mark.Delete(ctx, uint(markID)); err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}

func (h *Handler) FindPostsByUserID(c *gin.Context) error {
	ctx := c.Request.Context()
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return err
	}
	posts, err := h.services.Mark.FindPostsByUserID(ctx, uint(userID))
	if err != nil {
		return err
	}
	return schema.Respond(posts, c)
}
