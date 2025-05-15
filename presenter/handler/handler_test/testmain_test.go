package handler_test

import (
	"log"
	"os"
	"testing"

	"github.com/taimats/bhapi/testutils"
)

func TestMain(m *testing.M) {
	//テスト用に環境変数の設定
	if err := testutils.DotEnv(); err != nil {
		log.Fatalf("環境変数の設定に失敗:%v", err)
	}

	code := m.Run()

	os.Exit(code)
}
