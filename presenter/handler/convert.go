package handler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	ErrFailParse = errors.New("要素のパースに失敗")
)

// Json形式のUserをドメインのUser型に変換
func convertUser(u *User) (*domain.User, error) {
	var err error

	var id int64 = 0
	if u.Id != "" {
		id, err = strconv.ParseInt(u.Id, 10, 64)
		if err != nil {
			return nil, utils.NewErrChains(ErrFailParse, err)
		}
	}
	ca := time.Time{}
	if u.CreatedAt != "" {
		ca, err = parseStrTime(u.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	ua := time.Time{}
	if u.UpdatedAt != "" {
		ua, err = parseStrTime(u.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &domain.User{
		ID:         id,
		AuthUserId: u.AuthUserId,
		Name:       u.Name,
		Email:      domain.Email(u.Email),
		Password:   domain.Password(u.Password),
		CreatedAt:  ca,
		UpdatedAt:  ua,
	}, nil
}

// Json形式BookをドメインのBook型に変換
func convertBook(b *Book) (*domain.Book, error) {
	var err error

	var id int64 = 0
	if b.Id != "" {
		id, err = strconv.ParseInt(b.Id, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("idの数値変換に失敗:%w", err)
		}
	}
	price, err := strconv.Atoi(strings.ReplaceAll(b.Price, ",", ""))
	if err != nil {
		return nil, fmt.Errorf("priceの数値変換に失敗:%w", err)
	}
	page, err := strconv.Atoi(strings.ReplaceAll(b.Page, ",", ""))
	if err != nil {
		return nil, fmt.Errorf("pageの数値変換に失敗:%w", err)
	}
	ca := time.Now()
	if b.CreatedAt != "" {
		ca, err = parseStrTime(b.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	ua := time.Now()
	if b.UpdatedAt != "" {
		ca, err = parseStrTime(b.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	book := &domain.Book{
		ID:         id,
		ISBN10:     b.Isbn10,
		ImageURL:   b.ImageURL,
		Title:      b.Title,
		Author:     b.Author,
		Page:       page,
		Price:      price,
		BookStatus: domain.BookStatus(b.BookStatus),
		AuthUserId: b.AuthUserId,
		CreatedAt:  ca,
		UpdatedAt:  ua,
	}

	return book, nil
}

// ドメインUser型の配列をJson形式用に調整
func tweakUserForJSON(u *domain.User) *User {
	return &User{
		Id:         fmt.Sprint(u.ID),
		Name:       fmt.Sprint(u.Name),
		AuthUserId: fmt.Sprint(u.AuthUserId),
		Email:      fmt.Sprint(u.Email),
		Password:   u.Password.String(),
		CreatedAt:  u.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  u.UpdatedAt.Format(time.RFC3339),
	}
}

// ドメインBook型の配列をJson形式用に調整
func tweakBooksForJSON(books []*domain.Book) []*Book {
	if len(books) == 0 {
		return []*Book{}
	}

	//3桁カンマ区切りで出力するためのfmt拡張
	fmtx := message.NewPrinter(language.Japanese)

	updateBooks := make([]*Book, len(books))
	for i, book := range books {
		b := &Book{
			Id:         strconv.FormatInt(book.ID, 10),
			Isbn10:     book.ISBN10,
			ImageURL:   book.ImageURL,
			Title:      book.Title,
			Author:     book.Author,
			Page:       fmtx.Sprint(book.Page),
			Price:      fmtx.Sprint(book.Price),
			BookStatus: string(book.BookStatus),
			AuthUserId: book.AuthUserId,
			CreatedAt:  book.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  book.UpdatedAt.Format(time.RFC3339),
		}
		updateBooks[i] = b
	}

	return updateBooks
}

// ドメインRecord型の配列をJson形式に調整
func tweakRecordForJSON(dr *domain.Record) *Record {
	//3桁カンマ区切りで出力するためのfmt拡張
	fmtx := message.NewPrinter(language.Japanese)

	record := &Record{
		Costs:       fmtx.Sprint(dr.Costs),
		CostsRead:   fmtx.Sprint(dr.CostsRead),
		Volumes:     fmtx.Sprint(dr.Volumes),
		VolumesRead: fmtx.Sprint(dr.VolumesRead),
		Pages:       fmtx.Sprint(dr.Pages),
		PagesRead:   fmtx.Sprint(dr.PagesRead),
	}

	return record
}

// ドメインCharts型の配列をJson形式に調整
func tweakChartsForJSON(chs []*domain.Chart) []*Chart {
	charts := make([]*Chart, len(chs))

	for i, c := range chs {
		chart := &Chart{
			Label: fmt.Sprint(c.Label),
			Year:  fmt.Sprint(c.Year),
			Month: fmt.Sprintf("%v月", c.Month),
			Data:  fmt.Sprint(c.Data),
		}
		charts[i] = chart
	}

	return charts
}

// RFC3339形式のtime文字列をtime.Time型に変換するヘルパー関数。
// 引数sにはRFC3339形式(例."2006-01-02T15:04:05Z07:00")の文字列を入れる。
func parseStrTime(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, utils.NewErrChains(ErrFailParse, err)
	}
	return t, nil
}
