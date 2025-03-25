package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/testutils"
)

func TestGetSearch(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//外部APIテストサーバーの準備
	ts := testutils.PseudoGoogleBooksAPIServer(t)
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("GoogleBooksAPIテストサーバーのurlパースに失敗:%v", err)
	}
	testURL := u.JoinPath("books", "v1", "volumes").String()
	t.Setenv("GOOGL_BOOKS_API_URL", testURL)

	//assert時の期待データ
	books := []*domain.BookResult{
		{
			ISBN10:   "4167110121",
			ImageURL: "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:    "容疑者Xの献身",
			Author:   "東野圭吾",
			Page:     "0",
			Price:    "0",
		},
		{
			ISBN10:   "",
			ImageURL: "http://books.google.com/books/content?id=eNjdDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:    "容疑者Xの献身",
			Author:   "東野圭吾",
			Page:     "234",
			Price:    "770",
		},
		{
			ISBN10:   "",
			ImageURL: "http://books.google.com/books/content?id=hQDeDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:    "容疑者Xの献身　無料試し読み版",
			Author:   "東野圭吾",
			Page:     "49",
			Price:    "0",
		},
		{
			ISBN10:   "416711013X",
			ImageURL: "http://books.google.com/books/content?id=1LcsAwEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:    "ガリレオの苦悩",
			Author:   "東野圭吾",
			Page:     "0",
			Price:    "0",
		},
		{
			ISBN10:   "4167110083",
			ImageURL: "http://books.google.com/books/content?id=xdM9ywAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:    "予知夢",
			Author:   "東野圭吾",
			Page:     "0",
			Price:    "0",
		},
	}
	byteBooks, err := json.Marshal(books)
	if err != nil {
		t.Fatalf("json変換に失敗%s", err)
	}
	var expected bytes.Buffer
	expected.Write(byteBooks)
	expected.Write([]byte("\n"))

	//request, resposeの準備
	q := make(url.Values)
	q.Set("q", "容疑者の献身")
	r := httptest.NewRequest(http.MethodGet, "/search/?"+q.Encode(), nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)

	a := assert.New(t)

	//Act ***************
	err = sut.GetSearch(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.Equal(expected.Bytes(), w.Body.Bytes())
}
