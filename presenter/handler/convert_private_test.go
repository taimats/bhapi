package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/utils"
)

func TestTweakBooksForJSON(t *testing.T) {
	//Arrange
	cl := utils.NewTestClocker()

	books := []*domain.Book{
		{
			ID:         4167,
			ISBN10:     "4167110121",
			ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       2110,
			Price:      8800,
			BookStatus: domain.Bought,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
	}
	want := []*Book{
		{
			Id:         "4167",
			Isbn10:     "4167110121",
			ImageURL:   "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:      "容疑者Xの献身",
			Author:     "東野圭吾",
			Page:       "2,110",
			Price:      "8,800",
			BookStatus: "bought",
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			CreatedAt:  "2024-02-05 14:43:00",
			UpdatedAt:  "2024-02-05 14:43:00",
		},
	}

	//Act
	got := tweakBooksForJSON(books)

	//Assert
	assert.Equal(t, want, got)
}
