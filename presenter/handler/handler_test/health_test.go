package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/testutils"
)

func TestGetHealth(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)

	//assert時の評価データ
	expected := fmt.Sprintf("%s\n", `{"message":"ok"}`)

	a := assert.New(t)

	//Act ***************
	err := sut.GetHealth(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.Equal(expected, w.Body.String())
}

func TestGetHealthDb(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodGet, "/health/db", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)

	//assert時の評価データ
	expected := fmt.Sprintf("%s\n", `{"message":"ok"}`)

	a := assert.New(t)

	//Act ***************
	err := sut.GetHealth(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.Equal(expected, w.Body.String())
}
