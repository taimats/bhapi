package controller_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/testutils"
)

func TestSearchBooks(t *testing.T) {
	//Arrange ***************
	want := []*domain.BookResult{
		{
			ISBN10:   "4167110121",
			ImageURL: "http://books.google.com/books/content?id=TL3APAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:    "容疑者Xの献身",
			Author:   "東野圭吾",
			Page:     "0",
			Price:    "0",
		},
		{
			ISBN10:   "",
			ImageURL: "http://books.google.com/books/content?id=eNjdDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:    "容疑者Xの献身",
			Author:   "東野圭吾",
			Page:     "234",
			Price:    "770",
		},
		{
			ISBN10:   "",
			ImageURL: "http://books.google.com/books/content?id=hQDeDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
			Title:    "容疑者Xの献身　無料試し読み版",
			Author:   "東野圭吾",
			Page:     "49",
			Price:    "0",
		},
		{
			ISBN10:   "416711013X",
			ImageURL: "http://books.google.com/books/content?id=1LcsAwEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:    "ガリレオの苦悩",
			Author:   "東野圭吾",
			Page:     "0",
			Price:    "0",
		},
		{
			ISBN10:   "4167110083",
			ImageURL: "http://books.google.com/books/content?id=xdM9ywAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api",
			Title:    "予知夢",
			Author:   "東野圭吾",
			Page:     "0",
			Price:    "0",
		},
	}
	q := "容疑者の献身"

	ts := testutils.PseudoGoogleBooksAPIServer(t)
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("GoogleBooksAPIテストサーバーでurlパースに失敗:%v", err)
	}
	testURL := u.JoinPath("books", "v1", "volumes").String()

	ctx := context.Background()
	sut := controller.NewSearchBooks()

	a := assert.New(t)

	//Act ***************
	got, err := sut.SearchBooks(ctx, q, testURL)

	//Assert ***************
	a.Nil(err)
	a.Equal(want, got)
}
