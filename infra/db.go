package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewDBConfig() (dsn string) {
	u := os.Getenv("POSTGRES_USER")
	pd := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	h := os.Getenv("POSTGRES_HOST")
	pt := os.Getenv("POSTGRES_PORT")

	dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", u, pd, h, pt, db)

	return dsn
}

func NewDatabaseConnection(dsn string) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// データベース接続をテスト
	if err := sqldb.Ping(); err != nil {
		log.Fatalf("データベースの接続に失敗: %v", err)
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	// クエリーフックを追加することで、SQLを実行したクエリーが標準出力される
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	log.Println("データベースの接続に成功")

	return db, nil
}
