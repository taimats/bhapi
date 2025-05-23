package handler_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/testutils"
)

func TestGetHealth(t *testing.T) {
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

	sut, e := testutils.SetupHandler(bundb)
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	c, w := testutils.EchoContextWithRecorder(r, e)

	a := assert.New(t)
	g := goldie.New(t, goldie.WithDiffEngine(goldie.ColoredDiff))

	//Act ***************
	err = sut.GetHealth(c)
	resBody := testutils.IndentForJSON(t, w.Body.String())

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	g.Assert(t, t.Name(), resBody)
}

func TestGetHealthDb(t *testing.T) {
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

	sut, e := testutils.SetupHandler(bundb)
	r := httptest.NewRequest(http.MethodGet, "/health/db", nil)
	c, w := testutils.EchoContextWithRecorder(r, e)

	a := assert.New(t)
	g := goldie.New(t, goldie.WithDiffEngine(goldie.ColoredDiff))

	//Act ***************
	err = sut.GetHealthDb(c)
	resBody := testutils.IndentForJSON(t, w.Body.String())

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	g.Assert(t, t.Name(), resBody)
}
