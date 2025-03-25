package controller_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/testutils"
	"github.com/taimats/bhapi/utils"
)

func TestGetCharts(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	container, Terminate, err := testutils.SetUpDBContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { Terminate(container.Container) })

	err = testutils.MigrateUp(container.DSN, MigrationsPath)
	if err != nil {
		t.Fatalf("マイグレーションに失敗:%s", err)
	}

	db := testutils.SetUpDBForController(t, container.DSN)
	cl := utils.NewTestClocker()
	cr := repository.NewChart(db, cl)

	book := &domain.Book{
		ISBN10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "容疑者Xの献身",
		Author:     "東野圭吾",
		Page:       247,
		Price:      980,
		BookStatus: domain.Read,
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
	}
	charts := []*domain.Chart{
		{
			Label:      domain.ChartPrice,
			Year:       2025,
			Month:      2,
			Data:       980,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
		},
		{
			Label:      domain.ChartVolumes,
			Year:       2025,
			Month:      2,
			Data:       1,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
		},
		{
			Label:      domain.ChartPages,
			Year:       2025,
			Month:      2,
			Data:       247,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
		},
		{
			Label:      domain.ChartPrice,
			Year:       2025,
			Month:      2,
			Data:       980,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
		},
		{
			Label:      domain.ChartVolumes,
			Year:       2025,
			Month:      2,
			Data:       1,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
		},
		{
			Label:      domain.ChartPages,
			Year:       2025,
			Month:      2,
			Data:       247,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
		},
	}

	want := []*domain.Chart{
		{
			Label: domain.ChartPrice,
			Year:  2025,
			Month: 2,
			Data:  1960,
		},
		{
			Label: domain.ChartVolumes,
			Year:  2025,
			Month: 2,
			Data:  2,
		},
		{
			Label: domain.ChartPages,
			Year:  2025,
			Month: 2,
			Data:  494,
		},
	}

	testutils.InsertTestData(t, db, ctx, book)
	testutils.InsertTestData(t, db, ctx, charts...)

	sut := controller.NewChart(cr)
	authUserId := "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058"

	a := assert.New(t)

	//Act ***************
	got, err := sut.GetCharts(ctx, authUserId)

	//Act ***************
	a.Nil(err)
	a.Equal(want, got)
}
