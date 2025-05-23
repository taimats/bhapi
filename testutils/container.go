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

type DBContainer struct {
	Pctr    *postgres.PostgresContainer
	Dsn     string
	MigPath string
}

// テストケースごとに、DBコンテナをスナップショット状態に戻すテストヘルパー関数。
// リストアに失敗した場合、他のテストケースに影響があるため、リストアが失敗した時点でテストを終了。
// 内部でt.Cleanupしているため、呼び出すときはdeferやt.Cleanupは不要
func (dbctr *DBContainer) Restore(ctx context.Context, t *testing.T) {
	t.Cleanup(func() {
		if err := dbctr.Pctr.Restore(ctx); err != nil {
			t.Fatalf("DBコンテナのリストアに失敗:%s", err)
		}
		log.Println("DBコンテナをリストアしました")
	})
}

// パッケージ単位でテスト終了時にコンテナを破棄する
func (dbctr *DBContainer) Terminate() {
	if err := testcontainers.TerminateContainer(dbctr.Pctr); err != nil {
		log.Printf("コンテナの破棄に失敗:%s", err)
	}
}

// パッケージ単位で再利用可能なPostgres用コンテナの生成。内部でスナップショットをとり、
// テストケースごとにクリーンなDBサーバーを用意。
func SetupDBContainer(ctx context.Context, migPath string) (dbctr *DBContainer) {
	pctr, terminate, err := newDBContainer(ctx)
	if err != nil {
		log.Fatalf("DBコンテナの生成に失敗:%s", err)
	}
	dsn, err := pctr.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		terminate(pctr)
		log.Fatal(err)
	}
	//マイグレーション
	if err = migrateUp(dsn, migPath); err != nil {
		terminate(pctr)
		log.Fatalf("マイグレーションに失敗:%s", err)
	}
	//スナップショット
	if err := pctr.Snapshot(ctx); err != nil {
		terminate(pctr)
		log.Fatalf("スナップショットの生成に失敗:%s", err)
	}

	return &DBContainer{
		Pctr:    pctr,
		Dsn:     dsn,
		MigPath: migPath,
	}
}

func InsertTestData[T any](ctx context.Context, t *testing.T, db *bun.DB, data ...T) {
	t.Helper()

	err := db.NewInsert().Model(&data).Scan(ctx)
	if err != nil {
		t.Fatalf("テストデータの挿入に失敗:%s", err)
	}
}

func newDBContainer(ctx context.Context) (pctr *postgres.PostgresContainer, terminate func(c *postgres.PostgresContainer), err error) {
	//DB情報の取得
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	//DBコンテナの取得
	pctr, err = postgres.Run(ctx, "postgres:16-alpine",
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

	//内部で使用する後始末用の関数
	terminate = func(pctr *postgres.PostgresContainer) {
		if err := testcontainers.TerminateContainer(pctr); err != nil {
			log.Printf("コンテナの破棄に失敗:%s", err)
		}
	}
	return pctr, terminate, nil
}

func migrateUp(dsn string, migPath string) error {
	m, err := migrate.New("file://"+migPath, dsn)
	if err != nil {
		log.Fatal(err)
	}
	//マイグレーション時の接続を閉じないと
	//コンテナのスナップショットの生成に失敗する。
	defer func() {
		if srcErr, dbErr := m.Close(); srcErr != nil || dbErr != nil {
			log.Printf("srcErr:%v, dbErr:%v\n", srcErr, dbErr)
		}
	}()

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	return nil
}
