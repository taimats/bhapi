package handler_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/taimats/bhapi/testutils"
	"github.com/taimats/bhapi/utils"
)

var (
	dbctr *testutils.DBContainer
	cl    utils.TestClocker
)

// パッケージ内で共通する前後処理
// ・コンテナの生成
// ・コンテナの破棄
//
// ＜NOTE＞
// 各テストケースでは、dbctr.Restore()を実行すること。
// これにより、各テストケースでクリーンなDBが用意可能。
func TestMain(m *testing.M) {
	if err := testutils.DotEnv(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	migPath := "../../infra/migrations"
	dbctr = testutils.SetupDBContainer(ctx, migPath)
	cl = utils.NewTestClocker()

	code := m.Run()

	dbctr.Terminate()

	os.Exit(code)
}
