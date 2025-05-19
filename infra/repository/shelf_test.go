package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/testutils"
)

func TestFindBooksByAuthUserID(t *testing.T) {
	//Arrange
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	authUserId := "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058"

	books := []*domain.Book{
		{
			ID:         int64(1),
			ISBN10:     "4167110121",
			ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       330,
			Price:      1640,
			BookStatus: domain.Read,
			AuthUserId: authUserId,
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(2),
			ISBN10:     "",
			ImageURL:   "http://books.google.com/books/content?id=eNjdDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       234,
			Price:      770,
			BookStatus: domain.Reading,
			AuthUserId: authUserId,
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(3),
			ISBN10:     "",
			ImageURL:   "http://books.google.com/books/content?id=hQDeDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:      "容疑者Xの献身　無料試し読み版",
			Author:     "東野圭吾",
			Page:       49,
			Price:      0,
			BookStatus: domain.Bought,
			AuthUserId: authUserId,
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(4),
			ISBN10:     "416711013X",
			ImageURL:   "http://books.google.com/books/content?id=1LcsAwEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "ガリレオの苦悩",
			Author:     "東野圭吾",
			Page:       890,
			Price:      1240,
			BookStatus: domain.Read,
			AuthUserId: authUserId,
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(5),
			ISBN10:     "4167110083",
			ImageURL:   "http://books.google.com/books/content?id=xdM9ywAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "予知夢",
			Author:     "東野圭吾",
			Page:       220,
			Price:      220,
			BookStatus: domain.Bought,
			AuthUserId: authUserId,
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
	}
	testutils.InsertTestData(ctx, t, bundb, books...)
	sut := repository.NewShelf(bundb, cl)

	a := assert.New(t)

	//Act
	got, err := sut.FindBooksByAuthUserID(ctx, authUserId)

	//Assert
	a.Nil(err)
	a.Equal(books, got)
}

func TestCreateBookWithCharts(t *testing.T) {
	//Arrange
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
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
		},
		{
			Label:      domain.ChartVolumes,
			Year:       2025,
			Month:      2,
			Data:       1,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
		{
			Label:      domain.ChartPages,
			Year:       2025,
			Month:      2,
			Data:       247,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		},
	}
	sut := repository.NewShelf(bundb, cl)

	a := assert.New(t)

	//Act
	err = sut.CreateBookWithCharts(ctx, book, charts)

	//Assert
	a.Nil(err)
}

func TestUpdateBookWithCharts(t *testing.T) {
	//Arrange
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	book := &domain.Book{
		ID:         int64(1),
		ISBN10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "容疑者Xの献身",
		Author:     "東野圭吾",
		Page:       247,
		Price:      980,
		BookStatus: domain.Reading,
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		CreatedAt:  cl.Now(),
		UpdatedAt:  cl.Now(),
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
	}
	updatedBook := &domain.Book{
		ID:         int64(1),
		ISBN10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "容疑者Xの献身",
		Author:     "東野",
		Page:       300,
		Price:      1000,
		BookStatus: domain.Bought,
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		CreatedAt:  cl.Now(),
		UpdatedAt:  cl.Now(),
	}
	testutils.InsertTestData(ctx, t, bundb, book)
	testutils.InsertTestData(ctx, t, bundb, charts...)

	sut := repository.NewShelf(bundb, cl)
	a := assert.New(t)

	//Act
	err = sut.UpdateBookWithCharts(ctx, updatedBook)

	//Assert
	a.Nil(err)
}

func TestDleteBooksWithCharts(t *testing.T) {
	//Arrange
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	books := []*domain.Book{
		{
			ID:         int64(1),
			ISBN10:     "4167110121",
			ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       247,
			Price:      980,
			BookStatus: domain.Reading,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(2),
			ISBN10:     "416711013X",
			ImageURL:   "http://books.google.com/books/content?id=1LcsAwEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "ガリレオの苦悩",
			Author:     "東野圭吾",
			Page:       890,
			Price:      1240,
			BookStatus: domain.Read,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
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
			Data:       1240,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(2),
		},
		{
			Label:      domain.ChartVolumes,
			Year:       2025,
			Month:      2,
			Data:       1,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(2),
		},
		{
			Label:      domain.ChartPages,
			Year:       2025,
			Month:      2,
			Data:       890,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(2),
		},
	}
	testutils.InsertTestData(ctx, t, bundb, books...)
	testutils.InsertTestData(ctx, t, bundb, charts...)

	sut := repository.NewShelf(bundb, cl)

	a := assert.New(t)

	//Act
	err = sut.DleteBooksWithCharts(ctx, books)

	//Assert
	a.Nil(err)
}
