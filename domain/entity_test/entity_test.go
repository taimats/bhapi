package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/taimats/bhapi/domain"
)

func TestNewRecordFromBooks(t *testing.T) {
	//Arrange
	books := []*domain.Book{
		{
			ISBN10:     "4167110121",
			ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       330,
			Price:      1640,
			BookStatus: domain.Read,
		},
		{
			ISBN10:     "",
			ImageURL:   "http://books.google.com/books/content?id=eNjdDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       234,
			Price:      770,
			BookStatus: domain.Reading,
		},
		{
			ISBN10:     "",
			ImageURL:   "http://books.google.com/books/content?id=hQDeDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:      "容疑者Xの献身　無料試し読み版",
			Author:     "東野圭吾",
			Page:       49,
			Price:      0,
			BookStatus: domain.Bought,
		},
		{
			ISBN10:     "416711013X",
			ImageURL:   "http://books.google.com/books/content?id=1LcsAwEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "ガリレオの苦悩",
			Author:     "東野圭吾",
			Page:       890,
			Price:      1240,
			BookStatus: domain.Read,
		},
		{
			ISBN10:     "4167110083",
			ImageURL:   "http://books.google.com/books/content?id=xdM9ywAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "予知夢",
			Author:     "東野圭吾",
			Page:       220,
			Price:      220,
			BookStatus: domain.Bought,
		},
	}

	want := &domain.Record{
		Costs:       3870,
		CostsRead:   2880,
		Volumes:     5,
		VolumesRead: 2,
		Pages:       1723,
		PagesRead:   1220,
	}

	//Act
	got := domain.NewRecordFromBooks(books)

	//Assert
	assert.Equal(t, want, got)
}
