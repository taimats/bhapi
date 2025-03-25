package handler_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/presenter/handler"
	"github.com/taimats/bhapi/testutils"
	"github.com/taimats/bhapi/utils"
)

func TestPostShelfAuthUserId(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//リクエストボディの準備
	book := &handler.Book{
		Author:     "東野圭吾",
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		BookStatus: "read",
		Id:         "",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Isbn10:     "4167110121",
		Page:       "247",
		Price:      "980",
		Title:      "容疑者Xの献身",
	}
	body := testutils.ConvertForJSON(t, book)

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodPost, "/shelf/", &body)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)

	//Act ***************
	err := sut.PostShelfAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusCreated, w.Code)
	a.Empty(w.Body.Bytes())
}

func TestGetShelfWithAuthUserId(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//テストデータの挿入
	book := &domain.Book{
		ISBN10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "容疑者Xの献身",
		Author:     "東野圭吾",
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Page:       247,
		Price:      980,
		BookStatus: "read",
	}
	testutils.InsertTestData(t, htls.DB, htls.Ctx, book)

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodGet, "/shelf/", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)

	//Act ***************
	err := sut.GetShelfWithAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.NotEmpty(w.Body.Bytes())
}

func TestPutShelfWithAuthUserId(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//テストデータの挿入
	book := &domain.Book{
		ISBN10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "容疑者Xの献身",
		Author:     "東野圭吾",
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Page:       247,
		Price:      980,
		BookStatus: "read",
	}
	testutils.InsertTestData(t, htls.DB, htls.Ctx, book)

	//リクエストボディの準備
	cl := utils.NewTestClocker()
	body := &handler.Book{
		Author:     "update",
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		BookStatus: "read",
		CreatedAt:  cl.Now().String(),
		Id:         "1",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Isbn10:     "4167110121",
		Page:       "742",
		Price:      "980",
		Title:      "update",
	}
	jb := testutils.ConvertForJSON(t, body)

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodPut, "/shelf/", &jb)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)

	//Act ***************
	err := sut.PutShelfWithAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.NotEmpty(w.Body.Bytes())
}

func TestDeleteShelfWithAuthUserId(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//テストデータの挿入
	book := &domain.Book{
		ISBN10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "容疑者Xの献身",
		Author:     "東野圭吾",
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Page:       247,
		Price:      980,
		BookStatus: "read",
	}
	charts := []*domain.Chart{
		{
			Label:      domain.ChartPrice,
			Year:       2025,
			Month:      2,
			Data:       980,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
		},
		{
			Label:      domain.ChartVolumes,
			Year:       2025,
			Month:      2,
			Data:       1,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
		},
		{
			Label:      domain.ChartPages,
			Year:       2025,
			Month:      2,
			Data:       247,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
		},
	}
	testutils.InsertTestData(t, htls.DB, htls.Ctx, book)
	testutils.InsertTestData(t, htls.DB, htls.Ctx, charts...)

	//request, resposeの準備
	q := make(url.Values)
	q.Set("bookId", "1")
	r := httptest.NewRequest(http.MethodDelete, "/shelf/:authUserId/?"+q.Encode(), nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)

	//Act ***************
	err := sut.DeleteShelfWithAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusNoContent, w.Code)
	a.Empty(w.Body.Bytes())
}
