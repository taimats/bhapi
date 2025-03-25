package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/testutils"
)

func TestGetChartsWithAuthUserId(t *testing.T) {
	//Arrange ***************
	htls := testutils.PreSetUpForHandlerTest(t)
	t.Cleanup(func() { htls.Terminate(htls.Container.Container) })

	//handlerの準備
	sut, e := testutils.SetUpHandler(htls.DB)

	//テストデータの挿入
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
	testutils.InsertTestData(t, htls.DB, htls.Ctx, charts...)

	//request, resposeの準備
	r := httptest.NewRequest(http.MethodGet, "/charts", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetParamNames("authUserId")
	c.SetParamValues("c0cc3f0c-9a02-45ba-9de7-7d7276bb6058")

	a := assert.New(t)

	//Act ***************
	err := sut.GetChartsWithAuthUserId(c)

	//Assert ***************
	a.Nil(err)
	a.Equal(http.StatusOK, w.Code)
	a.NotEmpty(w.Body.Bytes())
}
