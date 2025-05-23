package domain

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type BookResult struct {
	ISBN10   string `json:"isbn10"`
	ImageURL string `json:"imageURL"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Page     string `json:"page"`
	Price    string `json:"price"`
}

const keyname = "GOOGLE_BOOKS_API_KEY"

// GoogleBooksAPI（外部サービス）に本の検索を実施
func SearchForGoogleBooks(query string, apiBaseURL string) ([]*BookResult, error) {
	//引数についてのvalidation
	if apiBaseURL == "" {
		return nil, errors.New("urlを入力ください")
	}

	//GoogleBooksAPIへのリクエスト処理
	apikey := os.Getenv(keyname)
	u, err := url.Parse(apiBaseURL)
	if err != nil {
		return nil, fmt.Errorf("urlのパースに失敗:%v", err)
	}
	q := u.Query()
	q.Set("q", query)
	q.Set("key", apikey)
	q.Set("startIndex", "0")
	q.Set("maxResults", "10")
	u.RawQuery = q.Encode()

	//レスポンスの処理
	res, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("リクエストに失敗:%w", err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	//res.bodyから必要な箇所のみ抽出
	jb, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("res.bodyの読み出しに失敗:%w", err)
	}

	//jsonから必要なものを抽出し、BookResult型の配列を作成
	books, err := ExtractBooksFromJSON(string(jb))
	if err != nil {
		return nil, fmt.Errorf("res.bodyの読み出しに失敗:%v", err)
	}

	return books, nil
}

// 外部APIからjsonを加工して、BookResult型の配列を生成
func ExtractBooksFromJSON(json string) ([]*BookResult, error) {
	//json形式になっているか確認
	if !gjson.Valid(json) {
		return nil, errors.New("invalid JSON")
	}
	//jsonの各項目に直接アクセスするためgjsonを利用
	js := gjson.Get(string(json), "items")

	books := make([]*BookResult, len(js.Array()))
	for i, j := range js.Array() {
		title := gjson.Get(j.String(), "volumeInfo.title").String()
		a := gjson.Get(j.String(), "volumeInfo.authors")
		var authors []string
		for _, author := range a.Array() {
			authors = append(authors, author.String())
		}
		isbn10 := gjson.Get(j.String(), `volumeInfo.industryIdentifiers.#(type="ISBN_10").identifier`).String()
		page := gjson.Get(j.String(), "volumeInfo.pageCount").Int()
		imageURL := gjson.Get(j.String(), "volumeInfo.imageLinks.thumbnail").String()
		price := gjson.Get(j.String(), "saleInfo.listPrice.amount").Int()

		//3桁カンマ区切りで出力するためのfmt拡張
		fmtx := message.NewPrinter(language.Japanese)

		b := &BookResult{
			ISBN10:   isbn10,
			ImageURL: imageURL,
			Title:    title,
			Author:   strings.Join(authors, "、"), //配列から[]を削除する処理
			Page:     fmtx.Sprint(page),          //数値をカンマ(,)区切りにする処理
			Price:    fmtx.Sprint(price),         //数値をカンマ(,)区切りにする処理
		}
		books[i] = b
	}

	return books, nil
}
