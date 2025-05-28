package handler_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/presenter/handler"
	"github.com/taimats/bhapi/testutils"
)

func TestPostShelfAuthUserId(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := bundb.Close(); err != nil {
			log.Println(err)
		}
	}()

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
	jb := testutils.ConvertToJSON(t, book)

	sut, e := testutils.SetupHandler(bundb)
	r := httptest.NewRequest(http.MethodPost, "/shelf/c0cc3f0c-9a02-45ba-9de7-7d7276bb6058", &jb)
	c, w := testutils.EchoContextWithRecorder(r, e)

	a := assert.New(t)

	//Act ***************
	err = sut.PostShelfAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusCreated, w.Code)
	a.Empty(w.Body.Bytes())
}

func TestGetShelfWithAuthUserId(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := bundb.Close(); err != nil {
			log.Println(err)
		}
	}()

	//テストデータの挿入
	book := &domain.Book{
		ID:         int64(1),
		ISBN10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "容疑者Xの献身",
		Author:     "東野圭吾",
		Page:       247,
		Price:      980,
		BookStatus: domain.Read,
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		CreatedAt:  cl.Now(),
		UpdatedAt:  cl.Now(),
	}
	testutils.InsertTestData(ctx, t, bundb, book)

	//request, resposeの準備
	sut, e := testutils.SetupHandler(bundb)
	r := httptest.NewRequest(http.MethodGet, "/shelf/:authUserId", nil)
	c, w := testutils.EchoContextWithRecorder(r, e)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)
	g := goldie.New(t, goldie.WithDiffEngine(goldie.ColoredDiff))

	//Act ***************
	err = sut.GetShelfWithAuthUserId(c)

	//Assert ***************
	resBody := testutils.IndentForJSON(t, w.Body.String())
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	g.Assert(t, t.Name(), resBody)
}

func TestPutShelfWithAuthUserId(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := bundb.Close(); err != nil {
			log.Println(err)
		}
	}()

	//テストデータの挿入
	book := &domain.Book{
		ID:         int64(1),
		ISBN10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "容疑者Xの献身",
		Author:     "東野圭吾",
		Page:       247,
		Price:      980,
		BookStatus: domain.Read,
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		CreatedAt:  cl.Now(),
		UpdatedAt:  cl.Now(),
	}
	charts := []*domain.Chart{
		{
			ID:         int64(1),
			Label:      domain.ChartPrice,
			Year:       2025,
			Month:      2,
			Data:       980,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(2),
			Label:      domain.ChartVolumes,
			Year:       2025,
			Month:      2,
			Data:       1,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(3),
			Label:      domain.ChartPages,
			Year:       2025,
			Month:      2,
			Data:       247,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
	}
	testutils.InsertTestData(ctx, t, bundb, book)
	testutils.InsertTestData(ctx, t, bundb, charts...)

	//リクエストボディの準備
	body := &handler.Book{
		Id:         "1",
		Isbn10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "update",
		Author:     "update",
		Page:       "742",
		Price:      "980",
		BookStatus: "read",
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		CreatedAt:  cl.NowString(),
		UpdatedAt:  cl.NowString(),
	}
	jb := testutils.ConvertToJSON(t, body)

	sut, e := testutils.SetupHandler(bundb)
	r := httptest.NewRequest(http.MethodPut, "/shelf/c0cc3f0c-9a02-45ba-9de7-7d7276bb6058", &jb)
	c, w := testutils.EchoContextWithRecorder(r, e)

	a := assert.New(t)

	//Act ***************
	err = sut.PutShelfWithAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.Empty(w.Body.Bytes())
}

func TestDeleteShelfWithAuthUserId(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := bundb.Close(); err != nil {
			log.Println(err)
		}
	}()

	//テストデータの挿入
	book := &domain.Book{
		ID:         int64(1),
		ISBN10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "容疑者Xの献身",
		Author:     "東野圭吾",
		Page:       247,
		Price:      980,
		BookStatus: domain.Read,
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		CreatedAt:  cl.Now(),
		UpdatedAt:  cl.Now(),
	}
	charts := []*domain.Chart{
		{
			ID:         int64(1),
			Label:      domain.ChartPrice,
			Year:       2025,
			Month:      2,
			Data:       980,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(2),
			Label:      domain.ChartVolumes,
			Year:       2025,
			Month:      2,
			Data:       1,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(3),
			Label:      domain.ChartPages,
			Year:       2025,
			Month:      2,
			Data:       247,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
	}
	testutils.InsertTestData(ctx, t, bundb, book)
	testutils.InsertTestData(ctx, t, bundb, charts...)

	sut, e := testutils.SetupHandler(bundb)
	q := make(url.Values)
	q.Set("bookId", "1")
	target := "/shelf/:authUserId?"
	r := httptest.NewRequest(http.MethodDelete, target+q.Encode(), nil)
	c, w := testutils.EchoContextWithRecorder(r, e)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)

	//Act ***************
	err = sut.DeleteShelfWithAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusNoContent, w.Code)
	a.Empty(w.Body.Bytes())
}
