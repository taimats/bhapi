package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/testutils"
)

func TestGetRecordsWithAuthUserId(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer bundb.Close()

	//テストデータの挿入
	books := []*domain.Book{
		{
			ID:         int64(1),
			ISBN10:     "4167110121",
			ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       330,
			Price:      1640,
			BookStatus: domain.Read,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(2),
			ISBN10:     "",
			ImageURL:   "http://books.google.com/books/content?id=eNjdDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       234,
			Price:      770,
			BookStatus: domain.Reading,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(3),
			ISBN10:     "",
			ImageURL:   "http://books.google.com/books/content?id=hQDeDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:      "容疑者Xの献身　無料試し読み版",
			Author:     "東野圭吾",
			Page:       49,
			Price:      0,
			BookStatus: domain.Bought,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(4),
			ISBN10:     "416711013X",
			ImageURL:   "http://books.google.com/books/content?id=1LcsAwEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "ガリレオの苦悩",
			Author:     "東野圭吾",
			Page:       890,
			Price:      1240,
			BookStatus: domain.Read,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(5),
			ISBN10:     "4167110083",
			ImageURL:   "http://books.google.com/books/content?id=xdM9ywAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "予知夢",
			Author:     "東野圭吾",
			Page:       220,
			Price:      220,
			BookStatus: domain.Bought,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
	}
	testutils.InsertTestData(ctx, t, bundb, books...)

	sut, e := testutils.SetupHandler(bundb)
	r := httptest.NewRequest(http.MethodGet, "/records/c0cc3f0c-9a02-45ba-9de7-7d7276bb6058", nil)
	c, w := testutils.EchoContextWithRecorder(r, e)

	a := assert.New(t)
	g := goldie.New(t, goldie.WithDiffEngine(goldie.ColoredDiff))

	//Act ***************
	err = sut.GetRecordsWithAuthUserId(c)

	//Assert ***************
	resBody := testutils.IndentForJSON(t, w.Body.String())
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	g.Assert(t, t.Name(), resBody)
}
