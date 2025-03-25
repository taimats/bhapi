package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/testutils"
	"github.com/taimats/bhapi/utils"
)

func TestFindBooksByAuthUserID(t *testing.T) {
	//Arrange
	db := testutils.SetUpDBForRepository(t)
	t.Cleanup(func() { db.Close() })

	authUserId := "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058"

	cl := utils.NewTestClocker()
	sr := repository.NewShelf(db, cl)

	ctx := context.Background()
	a := assert.New(t)

	//Act
	got, err := sr.FindBooksByAuthUserID(ctx, authUserId)

	//Assert
	a.Nil(err)
	a.NotNil(got)
}

func TestCreateBookWithCharts(t *testing.T) {
	//Arrange
	db := testutils.SetUpDBForRepository(t)
	t.Cleanup(func() { db.Close() })

	book := &domain.Book{
		ID:         0,
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

	cl := utils.NewTestClocker()
	sr := repository.NewShelf(db, cl)

	ctx := context.Background()
	a := assert.New(t)

	//Act
	err := sr.CreateBookWithCharts(ctx, book, charts)

	//Assert
	a.Nil(err)
}

func TestUpdateBookWithCharts(t *testing.T) {
	//Arrange
	db := testutils.SetUpDBForRepository(t)
	t.Cleanup(func() { db.Close() })

	cl := utils.NewTestClocker()
	book := &domain.Book{
		ID:         1,
		ISBN10:     "4167110121",
		ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
		Title:      "容疑者Xの献身",
		Author:     "東野圭吾",
		Page:       247,
		Price:      980,
		BookStatus: domain.Reading,
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		CreatedAt:  cl.Now(),
	}

	sr := repository.NewShelf(db, cl)

	ctx := context.Background()
	a := assert.New(t)

	//Act
	err := sr.UpdateBookWithCharts(ctx, book)

	//Assert
	a.Nil(err)
}

func TestDleteBooksWithCharts(t *testing.T) {
	//Arrange
	db := testutils.SetUpDBForRepository(t)
	t.Cleanup(func() { db.Close() })

	cl := utils.NewTestClocker()
	books := []*domain.Book{
		{
			ID:         int64(2),
			ISBN10:     "4167110121",
			ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       247,
			Price:      980,
			BookStatus: domain.Reading,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			CreatedAt:  cl.Now(),
		},
	}

	sr := repository.NewShelf(db, cl)

	ctx := context.Background()
	a := assert.New(t)

	//Act
	err := sr.DleteBooksWithCharts(ctx, books)

	//Assert
	a.Nil(err)
}
