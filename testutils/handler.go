package testutils

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/presenter/handler"
	"github.com/taimats/bhapi/utils"
	"github.com/uptrace/bun"
)

func SetUpHandler(db *bun.DB) (*handler.Handler, *echo.Echo) {
	//repositoryインスタンスの生成
	cl := utils.NewTestClocker()
	cr := repository.NewChart(db, cl)
	sr := repository.NewShelf(db, cl)
	ur := repository.NewUser(db, cl)

	//controllerインスタンスの生成
	cc := controller.NewChart(cr)
	sc := controller.NewShelf(sr)
	uc := controller.NewUser(ur)
	rc := controller.NewRecord(sr)
	sbc := controller.NewSearchBooks()

	e := echo.New()
	e.Validator = handler.NewCustomValidator(validator.New())

	//hanlderの設定
	server := handler.NewHandler(uc, cc, rc, sc, sbc)

	return server, e
}

func ConvertForJSON[T any](t *testing.T, data T) (js bytes.Buffer) {
	t.Helper()

	err := json.NewEncoder(&js).Encode(data)
	if err != nil {
		t.Fatalf("jsonのエンコードに失敗:%s", err)
	}

	return js
}
