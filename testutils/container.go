package testutils

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
)

type TerminateContainer func(pctr *postgres.PostgresContainer)

// パッケージ単位で再利用可能なPostgres用コンテナの生成。マイグレーション完了後にスナップショットをとる(pctr.Snapshot())ことで
// テストケースごとにクリーンなDBサーバーが用意可能。
func NewDBContainer(ctx context.Context) (pctr *postgres.PostgresContainer, Terminate func(c *postgres.PostgresContainer), err error) {
	//DB情報の取得
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	//DBコンテナの取得
	pctr, err = postgres.Run(ctx, "postgres:16",
		// postgres.WithInitScripts(scripts...),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
			wait.ForListeningPort("5432/tcp"),
		),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("DBコンテナの生成に失敗:%w", err)
	}

	//後始末用の関数
	Terminate = func(pctr *postgres.PostgresContainer) {
		if err := testcontainers.TerminateContainer(pctr); err != nil {
			log.Printf("コンテナの破棄に失敗:%s", err)
		}
	}

	return pctr, Terminate, nil
}

// テストケースごとにDBコンテナをスナップショット状態に戻す。
// リストアに失敗した場合、他のテストケースに影響があるため、リストアが失敗した時点でテストを終了。
func RestoreContainer(pctr *postgres.PostgresContainer, ctx context.Context, t *testing.T) {
	t.Cleanup(func() {
		if err := pctr.Restore(ctx); err != nil {
			t.Fatalf("DBコンテナのリストアに失敗:%s", err)
		}
	})
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

func InsertTestData[T any](ctx context.Context, t *testing.T, db *bun.DB, data ...T) {
	t.Helper()

	err := db.NewInsert().Model(&data).Scan(ctx)
	if err != nil {
		t.Fatalf("テストデータの挿入に失敗:%s", err)
	}
}
