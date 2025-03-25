package controller

import (
	"context"

	"github.com/taimats/bhapi/domain"
)

type SearchBooks struct{}

func NewSearchBooks() *SearchBooks {
	return &SearchBooks{}
}

func (sb *SearchBooks) SearchBooks(ctx context.Context, q string, baseURL string) ([]*domain.BookResult, error) {
	books, err := domain.SearchForGoogleBooks(q, baseURL)
	if err != nil {
		return nil, err
	}

	return books, nil
}
