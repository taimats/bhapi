package handler_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/testutils"
)

func TestGetSearch(t *testing.T) {
	//Arrange ***************
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer bundb.Close()

	//外部APIテストサーバーの準備
	ts := testutils.PseudoGoogleBooksAPIServer(t)
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("GoogleBooksAPIテストサーバーのurlパースに失敗:%v", err)
	}
	testURL := u.JoinPath("books", "v1", "volumes").String()
	t.Setenv("GOOGL_BOOKS_API_URL", testURL)

	sut, e := testutils.SetupHandler(bundb)
	q := make(url.Values)
	q.Set("q", "容疑者の献身")
	r := httptest.NewRequest(http.MethodGet, "/search?"+q.Encode(), nil)
	c, w := testutils.EchoContextWithRecorder(r, e)

	a := assert.New(t)
	g := goldie.New(t, goldie.WithDiffEngine(goldie.ColoredDiff))

	//Act ***************
	err = sut.GetSearch(c)

	//Assert ***************
	resBody := testutils.IndentForJSON(t, w.Body.String())
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	g.Assert(t, t.Name(), resBody)
}
