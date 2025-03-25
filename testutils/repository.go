package testutils

import (
	"testing"

	"github.com/taimats/bhapi/infra"
	"github.com/uptrace/bun"
)

func SetUpDBForController(t *testing.T, dsn string) *bun.DB {
	t.Helper()

	db, err := infra.NewDatabaseConnection(dsn)
	if err != nil {
		t.Fatalf("データベースの接続に失敗:%s", err)
	}

	return db
}

func SetUpDBForRepository(t *testing.T) *bun.DB {
	t.Helper()

	dsn := infra.NewDBConfig()
	db, err := infra.NewDatabaseConnection(dsn)
	if err != nil {
		t.Fatalf("データベースの接続に失敗:%s", err)
	}

	return db
}
