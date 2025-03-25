package controller

import (
	"context"

	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra/repository"
)

type Record struct {
	sr *repository.Shelf
}

func NewRecord(sr *repository.Shelf) *Record {
	return &Record{sr: sr}
}

func (rc *Record) GetRecord(ctx context.Context, authUserId string) (*domain.Record, error) {
	books, err := rc.sr.FindBooksByAuthUserID(ctx, authUserId)
	if err != nil {
		return nil, err
	}

	record := domain.NewRecordFromBooks(books)

	return record, nil
}
