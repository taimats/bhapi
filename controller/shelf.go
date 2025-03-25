package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra/repository"
)

type Shelf struct {
	sr *repository.Shelf
}

func NewShelf(sr *repository.Shelf) *Shelf {
	return &Shelf{sr: sr}
}

func (sc *Shelf) PostBookWithCharts(ctx context.Context, book *domain.Book) error {
	charts := domain.NewChartsFromBook(book)

	err := sc.sr.CreateBookWithCharts(ctx, book, charts)
	if err != nil {
		return err
	}

	return nil
}

func (sc *Shelf) GetShelf(ctx context.Context, authUserId string) ([]*domain.Book, error) {
	books, err := sc.sr.FindBooksByAuthUserID(ctx, authUserId)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (sc *Shelf) UpdateShelf(ctx context.Context, book *domain.Book) error {
	err := sc.sr.UpdateBookWithCharts(ctx, book)
	if err != nil {
		return err
	}

	return err
}

func (sc *Shelf) DeleteShelf(ctx context.Context, bookIds []string) error {
	books, err := newBooksFromBookIds(bookIds)
	if err != nil {
		return err
	}

	err = sc.sr.DleteBooksWithCharts(ctx, books)
	if err != nil {
		return err
	}

	return err
}

func newBooksFromBookIds(bookIds []string) ([]*domain.Book, error) {
	books := make([]*domain.Book, len(bookIds))

	for i, bi := range bookIds {
		id, err := strconv.ParseInt(bi, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("idの数値変換に失敗:%w", err)
		}
		b := &domain.Book{ID: id}
		books[i] = b
	}
	return books, nil
}
