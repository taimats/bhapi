package handler

import (
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/utils"
)

type Handler struct {
	uc  *controller.User
	cc  *controller.Chart
	rc  *controller.Record
	sc  *controller.Shelf
	sbc *controller.SearchBooks
	hc  *controller.HealthDB
}

func NewHandler(
	uc *controller.User,
	cc *controller.Chart,
	rc *controller.Record,
	sc *controller.Shelf,
	sbc *controller.SearchBooks,
	hc *controller.HealthDB,
) *Handler {
	return &Handler{
		uc:  uc,
		cc:  cc,
		rc:  rc,
		sc:  sc,
		sbc: sbc,
		hc:  hc,
	}
}

var _ HandlerInterface = (*Handler)(nil)

// user情報の登録
// (POST /auth/register)
func (h *Handler) PostAuthRegister(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストボディの読み込みに失敗")
	}

	if err := c.Validate(&u); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user, err := convertUser(&u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザー登録に失敗")
	}

	err = h.uc.RegisterUser(ctx, user)
	if err != nil {
		if errors.Is(err, utils.ErrAlrExists) {
			return echo.NewHTTPError(http.StatusBadRequest, "すでにユーザーが存在します")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザー登録に失敗")
	}

	return c.NoContent(http.StatusCreated)
}

// ユーザーごとにチャートデータを返す
// (GET /charts/{AuthUserId})
func (h *Handler) GetChartsWithAuthUserId(c echo.Context) error {
	authUserId := c.Param("authUserId")
	ctx := c.Request().Context()

	chs, err := h.cc.GetCharts(ctx, authUserId)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "図表がありません")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "図表の取得に失敗")
	}

	charts := tweakChartsForJSON(chs)
	return c.JSON(http.StatusOK, charts)
}

// サーバーの監視
// (GET /health)
func (h *Handler) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "ok",
	})
}

// DBサーバーの監視
// (GET /health/db)
func (h *Handler) GetHealthDb(c echo.Context) error {
	if h.hc.IsActive() {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "ok",
		})
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "DBに異常があります")
	}
}

// ユーザーごとに記録を返す
// (GET /records/{AuthUserId})
func (h *Handler) GetRecordsWithAuthUserId(c echo.Context) error {
	authUserId := c.Param("authUserId")
	ctx := c.Request().Context()

	record, err := h.rc.GetRecord(ctx, authUserId)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "記録がありません")
		}
		return echo.NewHTTPError(http.StatusNotFound, "記録の取得に失敗")
	}

	return c.JSON(http.StatusOK, tweakRecordForJSON(record))
}

// 書籍の検索結果を取得
// (GET /search)
func (h *Handler) GetSearch(c echo.Context) error {
	q := c.QueryParam("q")
	if q == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "検索文字を入力ください")
	}

	ctx := c.Request().Context()
	baseURL := os.Getenv("GOOGL_BOOKS_API_URL")

	results, err := h.sbc.SearchBooks(ctx, q, baseURL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "書籍の検索に失敗")
	}

	return c.JSON(http.StatusOK, results)
}

// ユーザーごとに本棚を複数削除
// (DELETE /shelf/{AuthUserId})
func (h *Handler) DeleteShelfWithAuthUserId(c echo.Context) error {
	bookIds := c.QueryParams()["bookId"]
	if len(bookIds) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "bookIdが必要です")
	}

	ctx := c.Request().Context()

	err := h.sc.DeleteShelf(ctx, bookIds)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "本の削除に失敗")
	}

	return c.NoContent(http.StatusNoContent)
}

// ユーザーごとに本棚を取得
// (GET /shelf/{AuthUserId})
func (h *Handler) GetShelfWithAuthUserId(c echo.Context) error {
	authUserId := c.Param("authUserId")
	ctx := c.Request().Context()

	books, err := h.sc.GetShelf(ctx, authUserId)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "本棚がありません")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "本棚の取得に失敗")
	}

	shelf := tweakBooksForJSON(books)

	return c.JSON(http.StatusOK, shelf)
}

// ユーザーごとに本を本棚に1冊ずつ作成
// (POST /shelf/{authUserId})
func (h *Handler) PostShelfAuthUserId(c echo.Context) error {
	var b Book
	if err := c.Bind(&b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストボディの取得に失敗")
	}

	if err := c.Validate(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "不正なリクエストです")
	}

	ctx := c.Request().Context()
	book, err := convertBook(&b)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "本の変換に失敗")
	}

	err = h.sc.PostBookWithCharts(ctx, book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "本の作成に失敗")
	}

	return c.NoContent(http.StatusCreated)
}

// ユーザーごとに本棚を1冊ずつ更新
// (PUT /shelf/{AuthUserId})
func (h *Handler) PutShelfWithAuthUserId(c echo.Context) error {
	b := new(Book)
	if err := c.Bind(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストボディの取得に失敗")
	}

	if err := c.Validate(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "不正なリクエストです")
	}

	book, err := convertBook(b)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "本の変換に失敗")
	}

	ctx := c.Request().Context()
	err = h.sc.UpdateShelf(ctx, book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "本の更新に失敗")
	}

	return c.NoContent(http.StatusOK)
}

// ユーザーを削除
// (DELETE /users/{AuthUserId})
func (h *Handler) DeleteUsersWithAuthUserId(c echo.Context) error {
	authUserId := c.Param("authUserId")
	ctx := c.Request().Context()

	err := h.uc.DeleteUser(ctx, authUserId)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "ユーザーがありません")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザーの削除に失敗")
	}

	return c.NoContent(http.StatusNoContent)
}

// ユーザー情報を返す
// (GET /users/{AuthUserId})
func (h *Handler) GetUsersWithAuthUserId(c echo.Context) error {
	authUserId := c.Param("authUserId")
	ctx := c.Request().Context()

	user, err := h.uc.GetUser(ctx, authUserId)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "ユーザーがありません")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザーの取得に失敗")
	}

	return c.JSON(http.StatusOK, tweakUserForJSON(user))
}

func (h *Handler) PutUsers(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストボディの読み込みに失敗")
	}
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "不正なリクエストです")
	}

	ctx := c.Request().Context()
	user, err := convertUser(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザーの変換に失敗")
	}

	err = h.uc.UpdateUser(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザーの更新に失敗")
	}

	return c.NoContent(http.StatusOK)
}
