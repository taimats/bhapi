package handler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/taimats/bhapi/domain"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func convertUser(u *User) *domain.User {
	const layout = "2006-01-02 15:04:05"
	ca, err := time.Parse(layout, u.CreatedAt)
	if err != nil {
		ca = time.Now()
	}

	return &domain.User{
		ID:         int64(0),
		AuthUserId: u.AuthUserId,
		Name:       u.Name,
		Email:      domain.Email(u.Email),
		Password:   domain.Password(u.Password),
		CreatedAt:  ca,
	}
}

func convertBook(b *Book) (*domain.Book, error) {
	price, err := strconv.Atoi(strings.Replace(b.Price, ",", "", -1))
	if err != nil {
		return nil, fmt.Errorf("priceの数値変換に失敗:%w", err)
	}

	page, err := strconv.Atoi(strings.Replace(b.Page, ",", "", -1))
	if err != nil {
		return nil, fmt.Errorf("pageの数値変換に失敗:%w", err)
	}

	var id int64
	if b.Id != "" {
		id, err = strconv.ParseInt(b.Id, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("idの数値変換に失敗:%w", err)
		}
	}

	const layout = "2006-01-02 15:04:05"
	ca, err := time.Parse(layout, b.CreatedAt)
	if err != nil {
		ca = time.Now()
	}
	ua, err := time.Parse(layout, b.UpdatedAt)
	if err != nil {
		ua = time.Now()
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

func tweakBooksForJSON(books []*domain.Book) []*Book {
	if len(books) == 0 {
		return []*Book{}
	}

	//3桁カンマ区切りで出力するためのfmt拡張
	fmtx := message.NewPrinter(language.Japanese)

	const layout = "2006-01-02 15:04:05"

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
			CreatedAt:  book.CreatedAt.Format(layout),
			UpdatedAt:  book.UpdatedAt.Format(layout),
		}

		updateBooks[i] = b
	}

	return updateBooks
}

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
