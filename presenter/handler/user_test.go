package handler_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/presenter/handler"
	"github.com/taimats/bhapi/testutils"
)

func TestPostAuthRegister(t *testing.T) {
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
	u := &handler.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Email:      "example@example.com",
	}
	jb := testutils.ConvertToJSON(t, u)

	//request, resposeの準備
	sut, e := testutils.SetupHandler(bundb)
	r := httptest.NewRequest(http.MethodPost, "/auth/register", &jb)
	c, w := testutils.EchoContextWithRecorder(r, e)

	a := assert.New(t)

	//Act ***************
	err = sut.PostAuthRegister(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusCreated, w.Code)
	a.Empty(w.Body.Bytes())
}

func TestGetUsersWithAuthUserId(t *testing.T) {
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
	u := &domain.User{
		ID:         int64(1),
		Name:       "example",
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Email:      "example@example.com",
		Password:   "7d7276bb6058",
		CreatedAt:  cl.NowJST(),
		UpdatedAt:  cl.NowJST(),
	}
	testutils.InsertTestData(ctx, t, bundb, u)

	sut, e := testutils.SetupHandler(bundb)
	r := httptest.NewRequest(http.MethodGet, "/users/:authUserId", nil)
	c, w := testutils.EchoContextWithRecorder(r, e)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)
	g := goldie.New(t, goldie.WithDiffEngine(goldie.ColoredDiff))

	//Act ***************
	err = sut.GetUsersWithAuthUserId(c)

	//Assert ***************
	resBody := testutils.IndentForJSON(t, w.Body.String())
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	g.Assert(t, t.Name(), resBody)
}

func TestDeleteUsersWithAuthUserId(t *testing.T) {
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
	u := &domain.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Email:      "example@example.com",
	}
	testutils.InsertTestData(ctx, t, bundb, u)

	sut, e := testutils.SetupHandler(bundb)
	r := httptest.NewRequest(http.MethodDelete, "/users/:authUserId", nil)
	c, w := testutils.EchoContextWithRecorder(r, e)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)

	//Act ***************
	err = sut.DeleteUsersWithAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusNoContent, w.Code)
	a.Empty(w.Body.Bytes())
}

func TestPutUsers(t *testing.T) {
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
	u := &domain.User{
		ID:         int64(1),
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Email:      "example@example.com",
		Password:   "",
		CreatedAt:  cl.Now(),
		UpdatedAt:  cl.Now(),
	}
	testutils.InsertTestData(ctx, t, bundb, u)

	//リクエストボディの準備
	updatedUser := &handler.User{
		Id:         "1",
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Email:      "example@example.com",
		Name:       "updated",
		CreatedAt:  cl.NowString(),
	}
	jb := testutils.ConvertToJSON(t, updatedUser)

	sut, e := testutils.SetupHandler(bundb)
	r := httptest.NewRequest(http.MethodPut, "/users", &jb)
	c, w := testutils.EchoContextWithRecorder(r, e)

	a := assert.New(t)

	//Act ***************
	err = sut.PutUsers(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.Empty(w.Body.Bytes())
}
