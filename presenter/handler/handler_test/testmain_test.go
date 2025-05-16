package handler_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/taimats/bhapi/testutils"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// テストケース全体で共有する
var pctr *postgres.PostgresContainer
var dsn string

func TestMain(m *testing.M) {
	//テスト用に環境変数の設定
	if err := testutils.DotEnv(); err != nil {
		log.Fatalf("環境変数の設定に失敗:%v", err)
	}

	//DBコンテナの生成
	ctx := context.Background()
	pctr, terminate, err := testutils.NewDBContainer(ctx)
	if err != nil {
		log.Fatalf("DBコンテナの生成に失敗:%s", err)
	}
	//マイグレーション
	dsn, err = pctr.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	migrationsPath := "../../infra/migrations"
	if err := testutils.MigrateUp(dsn, migrationsPath); err != nil {
		log.Fatalf("マイグレーションに失敗:%s", err)
	}
	//スナップショットでDBコンテナがリストア可能
	if err := pctr.Snapshot(ctx); err != nil {
		log.Fatalf("スナップショットの生成に失敗:%s", err)
	}

	code := m.Run()

	terminate(pctr)

	os.Exit(code)
}
