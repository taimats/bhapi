package testutils

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
)

type DBContainer struct {
	Container *postgres.PostgresContainer
	DSN       string
}

func SetUpDBContainer(ctx context.Context) (dbcontainer *DBContainer, Terminate func(c *postgres.PostgresContainer), err error) {
	//DB情報の取得
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	// //migration用のup.sqlファイル
	// scripts := readScripts()

	//DBコンテナの取得
	container, err := postgres.Run(ctx, "postgres:16",
		// postgres.WithInitScripts(scripts...),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForListeningPort("5432/tcp").
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("DBコンテナの生成に失敗:%w", err)
	}

	//あとでbun.DBの生成に必要なため取得
	dsn, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, nil, fmt.Errorf("DSNの取得に失敗:%w", err)
	}

	//後始末用の関数
	Terminate = func(c *postgres.PostgresContainer) {
		if err := testcontainers.TerminateContainer(c); err != nil {
			log.Printf("コンテナの破棄に失敗:%s\n", err)
		}
	}

	dbcontainer = &DBContainer{
		container,
		dsn,
	}

	return dbcontainer, Terminate, nil
}

func MigrateUp(dsn string, migrationsPath string) error {
	m, err := migrate.New("file://"+migrationsPath, dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func InsertTestData[T any](t *testing.T, db *bun.DB, ctx context.Context, data ...T) {
	t.Helper()

	err := db.NewInsert().Model(&data).Scan(ctx)
	if err != nil {
		t.Fatalf("テストデータの挿入に失敗:%s", err)
	}
}
