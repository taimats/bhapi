package handler

import (
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/utils"
)

type Server struct {
	uc  *controller.User
	cc  *controller.Chart
	rc  *controller.Record
	sc  *controller.Shelf
	sbc *controller.SearchBooks
}

func NewServer(
	uc *controller.User,
	cc *controller.Chart,
	rc *controller.Record,
	sc *controller.Shelf,
	sbc *controller.SearchBooks,
) *Server {
	return &Server{
		uc:  uc,
		cc:  cc,
		rc:  rc,
		sc:  sc,
		sbc: sbc,
	}
}

var _ ServerInterface = (*Server)(nil)

// user情報の登録
// (POST /auth/register)
func (s *Server) PostAuthRegister(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストボディの読み込みに失敗")
	}

	if err := c.Validate(&u); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user := convertUser(&u)

	err := s.uc.RegisterUser(ctx, user)
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
func (s *Server) GetChartsWithAuthUserId(c echo.Context) error {
	authUserId := c.Param("authUserId")
	ctx := c.Request().Context()

	chs, err := s.cc.GetCharts(ctx, authUserId)
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
func (s *Server) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "ok",
	})
}

// DBサーバーの監視
// (GET /health/db)
func (s *Server) GetHealthDb(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "ok",
	})
}

// ユーザーごとに記録を返す
// (GET /records/{AuthUserId})
func (s *Server) GetRecordsWithAuthUserId(c echo.Context) error {
	authUserId := c.Param("authUserId")
	ctx := c.Request().Context()

	record, err := s.rc.GetRecord(ctx, authUserId)
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
func (s *Server) GetSearch(c echo.Context) error {
	q := c.QueryParam("q")
	if q == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "検索文字を入力ください")
	}

	ctx := c.Request().Context()
	baseURL := os.Getenv("GOOGL_BOOKS_API_URL")

	results, err := s.sbc.SearchBooks(ctx, q, baseURL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "書籍の検索に失敗")
	}

	return c.JSON(http.StatusOK, results)
}

// ユーザーごとに本棚を複数削除
// (DELETE /shelf/{AuthUserId})
func (s *Server) DeleteShelfWithAuthUserId(c echo.Context) error {
	bookIds := c.QueryParams()["bookId"]
	if len(bookIds) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "bookIdが必要です")
	}

	ctx := c.Request().Context()

	err := s.sc.DeleteShelf(ctx, bookIds)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "本の削除に失敗")
	}

	return c.NoContent(http.StatusNoContent)
}

// ユーザーごとに本棚を取得
// (GET /shelf/{AuthUserId})
func (s *Server) GetShelfWithAuthUserId(c echo.Context) error {
	authUserId := c.Param("authUserId")
	ctx := c.Request().Context()

	books, err := s.sc.GetShelf(ctx, authUserId)
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
func (s *Server) PostShelfAuthUserId(c echo.Context) error {
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

	err = s.sc.PostBookWithCharts(ctx, book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "本の作成に失敗")
	}

	return c.NoContent(http.StatusCreated)
}

// ユーザーごとに本棚を1冊ずつ更新
// (PUT /shelf/{AuthUserId})
func (s *Server) PutShelfWithAuthUserId(c echo.Context) error {
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
	err = s.sc.UpdateShelf(ctx, book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "本の更新に失敗")
	}

	return c.NoContent(http.StatusOK)
}

// ユーザーを削除
// (DELETE /users/{AuthUserId})
func (s *Server) DeleteUsersWithAuthUserId(c echo.Context) error {
	authUserId := c.Param("authUserId")
	ctx := c.Request().Context()

	err := s.uc.DeleteUser(ctx, authUserId)
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
func (s *Server) GetUsersWithAuthUserId(c echo.Context) error {
	authUserId := c.Param("authUserId")
	ctx := c.Request().Context()

	user, err := s.uc.GetUser(ctx, authUserId)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "ユーザーがありません")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザーの取得に失敗")
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) PutUsers(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストボディの読み込みに失敗")
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user := convertUser(u)

	err := s.uc.UpdateUser(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザーの更新に失敗")
	}

	return c.NoContent(http.StatusOK)
}
