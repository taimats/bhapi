package controller_test

import (
	"context"
	"testing"

	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/infra"
)

func TestHealthDB(t *testing.T) {
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer bundb.Close()
	sut := controller.NewHealthDB(bundb)

	ok := sut.IsActive()

	if ok != true {
		t.Error("DBが正常ではありません")
	}
}
