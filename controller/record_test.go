package controller_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/testutils"
)

func TestGetRecord(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer bundb.Close()

	books := []*domain.Book{
		{
			ISBN10:     "4167110121",
			ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       330,
			Price:      1640,
			BookStatus: domain.Read,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
		{
			ISBN10:     "",
			ImageURL:   "http://books.google.com/books/content?id=eNjdDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       234,
			Price:      770,
			BookStatus: domain.Reading,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
		{
			ISBN10:     "",
			ImageURL:   "http://books.google.com/books/content?id=hQDeDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:      "容疑者Xの献身　無料試し読み版",
			Author:     "東野圭吾",
			Page:       49,
			Price:      0,
			BookStatus: domain.Bought,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
		{
			ISBN10:     "416711013X",
			ImageURL:   "http://books.google.com/books/content?id=1LcsAwEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "ガリレオの苦悩",
			Author:     "東野圭吾",
			Page:       890,
			Price:      1240,
			BookStatus: domain.Read,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
		{
			ISBN10:     "4167110083",
			ImageURL:   "http://books.google.com/books/content?id=xdM9ywAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "予知夢",
			Author:     "東野圭吾",
			Page:       220,
			Price:      220,
			BookStatus: domain.Bought,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
	}
	testutils.InsertTestData(ctx, t, bundb, books...)

	want := &domain.Record{
		Costs:       3870,
		CostsRead:   2880,
		Volumes:     5,
		VolumesRead: 2,
		Pages:       1723,
		PagesRead:   1220,
	}

	sr := repository.NewShelf(bundb, cl)
	sut := controller.NewRecord(sr)
	authUserId := "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058"

	a := assert.New(t)

	//Act ***************
	got, err := sut.GetRecord(ctx, authUserId)

	//Assert ***************
	a.Nil(err)
	a.Equal(want, got)
}
