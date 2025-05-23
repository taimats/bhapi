package controller_test

import (
	"context"
	"log"
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
	defer func() {
		if err := bundb.Close(); err != nil {
			log.Println(err)
		}
	}()
	sut := controller.NewHealthDB(bundb)

	ok := sut.IsActive()

	if ok != true {
		t.Error("DBが正常ではありません")
	}
}
