package testutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/presenter/handler"
	"github.com/taimats/bhapi/utils"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/uptrace/bun"
)

func SetUpHandler(db *bun.DB) (*handler.Server, *echo.Echo) {
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
	server := handler.NewServer(uc, cc, rc, sc, sbc)

	return server, e
}

func SetUpDBForHandler(dsn string) (*bun.DB, error) {
	db, err := infra.NewDatabaseConnection(dsn)
	if err != nil {
		return nil, fmt.Errorf("データベースの接続に失敗:%s", err)
	}

	return db, nil
}

func InsertTestDataForHandler[T any](db *bun.DB, ctx context.Context, data ...T) error {
	err := db.NewInsert().Model(&data).Scan(ctx)
	if err != nil {
		return fmt.Errorf("データの挿入に失敗:%w", err)
	}

	return nil
}

func ConvertForJSON[T any](t *testing.T, data T) (js bytes.Buffer) {
	t.Helper()

	err := json.NewEncoder(&js).Encode(data)
	if err != nil {
		t.Fatalf("jsonのエンコードに失敗:%s", err)
	}

	return js
}

type hanlderTestTools struct {
	DB        *bun.DB
	Ctx       context.Context
	Container *DBContainer
	Terminate func(c *postgres.PostgresContainer)
}

func PreSetUpForHandlerTest(t *testing.T) *hanlderTestTools {
	t.Helper()

	const MigrationsPath = "../../../infra/migrations"

	//DBコンテナの準備
	ctx := context.Background()
	container, Terminate, err := SetUpDBContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	//マイグレーションアップ
	err = MigrateUp(container.DSN, MigrationsPath)
	if err != nil {
		t.Fatalf("マイグレーションに失敗:%s", err)
	}

	//DBへの接続
	db, err := SetUpDBForHandler(container.DSN)
	if err != nil {
		t.Fatalf("DBの接続に失敗:%s", err)
	}

	return &hanlderTestTools{db, ctx, container, Terminate}
}
