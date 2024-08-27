package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initPost(v1 *gin.RouterGroup) {
	v1.GET(
		"/post",
		middleware.GinErrorHandle(h.GetListingPosts),
	)
}

func (h *Handler) GetListingPosts(c *gin.Context) error {
	ctx := c.Request.Context()
	var userId *string
	token := c.GetHeader("Authorization")
	if token != "" {
		userIdStr, err := h.services.Auth.VerifyToken(token)
		if err != nil {
			return err
		}
		userId = &userIdStr
	}

	posts, err := h.usecases.Post.GetListing(ctx, userId)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)

}
