package domain_test

import (
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/testutils"
)

func TestSearchForGoogleBooks(t *testing.T) {
	t.Parallel()
	//Arrange
	err := testutils.DotEnv()
	if err != nil {
		t.Fatal(err)
	}
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

	//Act
	got, err := domain.SearchForGoogleBooks(q, testURL)
	if err != nil {
		t.Fatal(err)
	}

	//Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("a mismatch between (-want and +got):%s", diff)
	}
}

func TestExtractBooksFromJSON(t *testing.T) {
	t.Parallel()
	//Arrange
	searchResult, err := testutils.TestFile("response_body.json")
	if err != nil {
		t.Fatalf("テストデータの取得に失敗:%s", err)
	}
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

	//Act
	got, err := domain.ExtractBooksFromJSON(string(searchResult))
	if err != nil {
		t.Fatal(err)
	}

	//Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("a mismatch between (-want +got):%s", diff)
	}
}
