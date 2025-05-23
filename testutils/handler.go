package testutils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/presenter/handler"
	"github.com/taimats/bhapi/utils"
	"github.com/uptrace/bun"
)

// テスト用のハンドラーとバリデーション登録済みのechoインスタンスを返す。
func SetupHandler(db *bun.DB) (*handler.Handler, *echo.Echo) {
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
	hc := controller.NewHealthDB(db)

	e := echo.New()
	e.Validator = handler.NewCustomValidator(validator.New())

	//hanlderの設定
	h := handler.NewHandler(uc, cc, rc, sc, sbc, hc)

	return h, e
}

// handlerに渡すテスト用のecho contextとresponseRecorderを返す
func EchoContextWithRecorder(r *http.Request, e *echo.Echo) (c echo.Context, w *httptest.ResponseRecorder) {
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w = httptest.NewRecorder()
	c = e.NewContext(r, w)
	return c, w
}

// Goの構造体をjsonに変換し、ioのやりとりができるようにbufferで返す
func ConvertToJSON[T any](t *testing.T, data T) bytes.Buffer {
	t.Helper()

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		t.Fatalf("jsonのエンコードに失敗:%s", err)
	}

	return buf
}

// リテラルな文字列をjson形式で表示
func IndentForJSON(t *testing.T, data string) []byte {
	t.Helper()

	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(data), "", "  ")
	if err != nil {
		t.Fatalf("jsonのインデントに失敗:%s", err)
	}

	return buf.Bytes()
}
