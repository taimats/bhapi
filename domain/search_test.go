package domain_test

import (
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
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
	ts := testutils.PseudoGoogleBooksAPIServer(t)
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("GoogleBooksAPIテストサーバーでurlパースに失敗:%v", err)
	}
	testURL := u.JoinPath("books", "v1", "volumes").String()
	q := "容疑者の献身"

	a := assert.New(t)
	g := goldie.New(t, goldie.WithDiffEngine(goldie.ColoredDiff))

	//Act
	results, err := domain.SearchForGoogleBooks(q, testURL)

	//Assert
	j := testutils.ConvertToJSON(t, results)
	got := testutils.IndentForJSON(t, j.String())
	a.Nil(err)
	g.Assert(t, t.Name(), got)
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
		t.Errorf("異なる構造体(-want +got):%s", diff)
	}
}

func BenchmarkSearchForGoogleBooks(b *testing.B) {
	err := testutils.DotEnv()
	if err != nil {
		b.Fatal(err)
	}
	ts := testutils.PseudoAPIServer(b)
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		b.Fatalf("GoogleBooksAPIテストサーバーでurlパースに失敗:%v", err)
	}
	testURL := u.JoinPath("books", "v1", "volumes").String()
	q := "容疑者の献身"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = domain.SearchForGoogleBooks(q, testURL)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkExtractBooksFromJSON(b *testing.B) {
	searchResult, err := testutils.TestFile("response_body.json")
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := domain.ExtractBooksFromJSON(string(searchResult))
		if err != nil {
			b.Fatal(err)
		}
	}
}
