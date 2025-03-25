package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/presenter/handler"
	"github.com/taimats/bhapi/testutils"
)

func TestPostAuthRegister(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//リクエストボディの準備
	u := &handler.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Email:      "example@example.com",
	}
	body := testutils.ConvertForJSON(t, u)

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodPost, "/auth/register", &body)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)

	a := assert.New(t)

	//Act ***************
	err := sut.PostAuthRegister(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusCreated, w.Code)
	a.Empty(w.Body.Bytes())
}

func TestGetUsersWithAuthUserId(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//テストデータの挿入
	u := &domain.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Email:      "example@example.com",
	}
	testutils.InsertTestData(t, htls.DB, htls.Ctx, u)

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodGet, "/users/", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)

	//Act ***************
	err := sut.GetUsersWithAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.NotEmpty(w.Body.Bytes())
}

func TestDeleteUsersWithAuthUserId(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//テストデータの挿入
	u := &domain.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Email:      "example@example.com",
	}
	testutils.InsertTestData(t, htls.DB, htls.Ctx, u)

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodDelete, "/users/", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)

	//Act ***************
	err := sut.DeleteUsersWithAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusNoContent, w.Code)
	a.Empty(w.Body.Bytes())
}

func TestPutUsers(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//テストデータの挿入
	u := &domain.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Email:      "example@example.com",
	}
	testutils.InsertTestData(t, htls.DB, htls.Ctx, u)

	//リクエストボディの準備
	uu := &handler.User{
		Id:         "1",
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Email:      "example@example.com",
		Name:       "updated",
	}
	body := testutils.ConvertForJSON(t, uu)

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodPut, "/users", &body)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)

	a := assert.New(t)

	//Act ***************
	err := sut.PutUsers(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.Empty(w.Body.String())
}
