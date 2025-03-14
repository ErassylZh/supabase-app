package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"work-project/internal/admin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initContest(v1 *gin.RouterGroup) {
	v1.GET(
		"/contest/active",
		middleware.GinErrorHandle(h.GetActiveContest),
	)
	v1.GET(
		"/contest",
		middleware.GinErrorHandle(h.GetContestData),
	)
	v1.POST(
		"/contest/join",
		middleware.GinErrorHandle(h.JoinContest),
	)
	v1.POST(
		"/contest/read",
		middleware.GinErrorHandle(h.ReadContestBook),
	)
	v1.GET(
		"/contest/prize",
		middleware.GinErrorHandle(h.GetContestPrizes),
	)
	v1.GET(
		"/contest/book",
		middleware.GinErrorHandle(h.GetContestBooks),
	)
	v1.GET(
		"/contest/book/by-id",
		middleware.GinErrorHandle(h.GetContestBookByID),
	)
	v1.POST(
		"/contest",
		middleware.GinErrorHandle(h.CreateContest),
	)
	v1.DELETE(
		"/contest",
		middleware.GinErrorHandle(h.DeleteContest),
	)
	v1.PUT(
		"/contest",
		middleware.GinErrorHandle(h.UpdateContest),
	)
	v1.GET(
		"/contest/all",
		middleware.GinErrorHandle(h.GetContestList),
	)
	v1.POST(
		"/contest/prize",
		middleware.GinErrorHandle(h.CreateContestPrizes),
	)
	v1.POST(
		"/contest/book",
		middleware.GinErrorHandle(h.CreateContestBooks),
	)
	v1.DELETE(
		"/contest/prize",
		middleware.GinErrorHandle(h.DeleteContestPrizes),
	)
	v1.DELETE(
		"/contest/book",
		middleware.GinErrorHandle(h.DeleteContestBooks),
	)
	v1.PUT(
		"/contest/prize",
		middleware.GinErrorHandle(h.UpdateContestPrizes),
	)
	v1.PUT(
		"/contest/book",
		middleware.GinErrorHandle(h.UpdateContestBooks),
	)
}

// GetContestData
// WhoAmi godoc
// @Summary получить активные контесты
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]schema.ContestData]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param contest_id query int true "contest_id"
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest [get]
func (h *Handler) GetContestData(c *gin.Context) error {
	ctx := c.Request.Context()

	var data schema.ContestGetRequest
	if err := c.BindQuery(&data); err != nil {
		return err
	}

	token := c.GetHeader("Authorization")
	userID, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	data.UserId = userID

	contest, err := h.services.Contest.Get(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(contest, c)
}

// GetActiveContest
// WhoAmi godoc
// @Summary получить активные контесты
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[schema.ContestActivity]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/active [get]
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
// @Success 200 {object} schema.Response[schema.Empty]
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

// ReadContestBook
// WhoAmi godoc
// @Summary Прочесть контест книгу
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[schema.ReadContestRequest]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param data body schema.ReadContestRequest true "CreateMark"
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/read [post]
func (h *Handler) ReadContestBook(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userID, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}

	var data schema.ReadContestRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		return err
	}
	data.UserID = userID

	res, err := h.services.Contest.Read(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(res, c)
}

// GetContestPrizes
// WhoAmi godoc
// @Summary получить призы контеста
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.ContestPrize]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param contest_id query int true "contest_id"
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/prize [get]
func (h *Handler) GetContestPrizes(c *gin.Context) error {
	ctx := c.Request.Context()

	contestId, err := strconv.ParseUint(c.Query("contest_id"), 10, 64)
	if err != nil {
		return err
	}

	contestPrizes, err := h.services.Contest.GetPrizes(ctx, uint(contestId))
	if err != nil {
		return err
	}
	return schema.Respond(contestPrizes, c)
}

// GetContestBooks
// WhoAmi godoc
// @Summary получить книги контеста
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.ContestBook]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param contest_id query int true "contest_id"
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/book [get]
func (h *Handler) GetContestBooks(c *gin.Context) error {
	ctx := c.Request.Context()

	contestId, err := strconv.ParseUint(c.Query("contest_id"), 10, 64)
	if err != nil {
		return err
	}

	contestPrizes, err := h.services.Contest.GetBooks(ctx, uint(contestId))
	if err != nil {
		return err
	}
	return schema.Respond(contestPrizes, c)
}

// GetContestBookByID
// WhoAmi godoc
// @Summary получить книги контеста
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.ContestBook]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param contest_book_id query int true "contest_book_id"
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/book/by-id [get]
func (h *Handler) GetContestBookByID(c *gin.Context) error {
	ctx := c.Request.Context()

	contestBookId, err := strconv.ParseUint(c.Query("contest_book_id"), 10, 64)
	if err != nil {
		return err
	}

	contestPrizes, err := h.services.Contest.GetBookByID(ctx, uint(contestBookId))
	if err != nil {
		return err
	}
	return schema.Respond(contestPrizes, c)
}

// CreateContest
// WhoAmi godoc
// @Summary создать contest
// @Accept json
// @Produce json
// @Param data body admin.CreateContest true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Contest]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest [post]
func (h *Handler) CreateContest(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreateContest
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	contest, err := h.services.Contest.Create(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(contest, c)
}

// UpdateContest
// WhoAmi godoc
// @Summary создать contest
// @Accept json
// @Produce json
// @Param data body admin.UpdateContest true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Contest]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest [put]
func (h *Handler) UpdateContest(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdateContest
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	contest, err := h.services.Contest.Update(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(contest, c)
}

// DeleteContest
// WhoAmi godoc
// @Summary удолит продукт
// @Accept json
// @Produce json
// @Param contest_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest [delete]
func (h *Handler) DeleteContest(c *gin.Context) error {
	ctx := c.Request.Context()

	productId, err := strconv.ParseUint(c.Query("contest_id"), 10, 64)
	if err != nil {
		return err
	}

	err = h.services.Contest.Delete(ctx, uint(productId))
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}

// GetContestList
// WhoAmi godoc
// @Summary удолит продукт
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.Contest]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/all [get]
func (h *Handler) GetContestList(c *gin.Context) error {
	ctx := c.Request.Context()

	contests, err := h.services.Contest.GetAll(ctx)
	if err != nil {
		return err
	}

	return schema.Respond(contests, c)
}

// CreateContestBooks
// WhoAmi godoc
// @Summary создать contest
// @Accept json
// @Produce json
// @Param data body admin.CreateContestBook true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Contest]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/book [post]
func (h *Handler) CreateContestBooks(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreateContestBook
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	contest, err := h.services.Contest.CreateBook(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(contest, c)
}

// UpdateContestBooks
// WhoAmi godoc
// @Summary создать contest
// @Accept json
// @Produce json
// @Param data body admin.UpdateContestBook true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Contest]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/book [put]
func (h *Handler) UpdateContestBooks(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdateContestBook
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	contest, err := h.services.Contest.UpdateBook(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(contest, c)
}

// DeleteContestBooks
// WhoAmi godoc
// @Summary удолит продукт
// @Accept json
// @Produce json
// @Param contest_book_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/books [delete]
func (h *Handler) DeleteContestBooks(c *gin.Context) error {
	ctx := c.Request.Context()

	productId, err := strconv.ParseUint(c.Query("contest_book_id"), 10, 64)
	if err != nil {
		return err
	}

	err = h.services.Contest.DeleteBook(ctx, uint(productId))
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}

// CreateContestPrizes
// WhoAmi godoc
// @Summary создать contest
// @Accept json
// @Produce json
// @Param data body admin.CreateContestPrize true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Contest]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/prize [post]
func (h *Handler) CreateContestPrizes(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreateContestPrize
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	contestPrize, err := h.services.Contest.CreatePrize(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(contestPrize, c)
}

// UpdateContestPrizes
// WhoAmi godoc
// @Summary создать contest
// @Accept json
// @Produce json
// @Param data body admin.UpdateContestPrize true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Contest]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/prize [put]
func (h *Handler) UpdateContestPrizes(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdateContestPrize
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	contest, err := h.services.Contest.UpdatePrize(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(contest, c)
}

// DeleteContestPrizes
// WhoAmi godoc
// @Summary удолит продукт
// @Accept json
// @Produce json
// @Param contest_prize_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags contest
// @Router /api/v1/contest/prize [delete]
func (h *Handler) DeleteContestPrizes(c *gin.Context) error {
	ctx := c.Request.Context()

	prizeId, err := strconv.ParseUint(c.Query("contest_prize_id"), 10, 64)
	if err != nil {
		return err
	}

	err = h.services.Contest.DeletePrize(ctx, uint(prizeId))
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}
