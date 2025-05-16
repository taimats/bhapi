package handler_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/taimats/bhapi/testutils"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var pctr *postgres.PostgresContainer
var dsn string

func TestMain(m *testing.M) {
	//テスト用に環境変数の設定
	if err := testutils.DotEnv(); err != nil {
		log.Fatalf("環境変数の設定に失敗:%v", err)
	}

	//DBコンテナの生成
	ctx := context.Background()
	migrationsPaths := "../../infra/migrations"
	ctr, name, terminate, err := testutils.SetUpDBContainer(ctx, migrationsPaths)
	if err != nil {
		log.Fatal(err)
	}
	//varで定義された変数に詰めなおしてパッケージ内で共有
	pctr = ctr
	dsn = name

	code := m.Run()

	terminate(pctr)

	os.Exit(code)
}
