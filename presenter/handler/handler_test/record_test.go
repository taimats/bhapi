package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/testutils"
)

func TestGetRecordsWithAuthUserId(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//テストデータの挿入
	books := []*domain.Book{
		{
			ISBN10:     "4167110121",
			ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       330,
			Price:      1640,
			BookStatus: domain.Read,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
		{
			ISBN10:     "",
			ImageURL:   "http://books.google.com/books/content?id=eNjdDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       234,
			Price:      770,
			BookStatus: domain.Reading,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
		{
			ISBN10:     "",
			ImageURL:   "http://books.google.com/books/content?id=hQDeDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:      "容疑者Xの献身　無料試し読み版",
			Author:     "東野圭吾",
			Page:       49,
			Price:      0,
			BookStatus: domain.Bought,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
		{
			ISBN10:     "416711013X",
			ImageURL:   "http://books.google.com/books/content?id=1LcsAwEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "ガリレオの苦悩",
			Author:     "東野圭吾",
			Page:       890,
			Price:      1240,
			BookStatus: domain.Read,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
		{
			ISBN10:     "4167110083",
			ImageURL:   "http://books.google.com/books/content?id=xdM9ywAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "予知夢",
			Author:     "東野圭吾",
			Page:       220,
			Price:      220,
			BookStatus: domain.Bought,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
	}

	testutils.InsertTestData(t, htls.DB, htls.Ctx, books...)

	record := &domain.Record{
		Costs:       3870,
		CostsRead:   2880,
		Volumes:     5,
		VolumesRead: 2,
		Pages:       1723,
		PagesRead:   1220,
	}
	byteRecord, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("json変換に失敗%s", err)
	}
	var expected bytes.Buffer
	expected.Write(byteRecord)
	expected.Write([]byte("\n"))

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodGet, "/records", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)

	//Act ***************
	err = sut.GetRecordsWithAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.Equal(expected.Bytes(), w.Body.Bytes())
}
