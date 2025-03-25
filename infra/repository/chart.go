package repository

import (
	"context"

	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/utils"
	"github.com/uptrace/bun"
)

type Chart struct {
	db *bun.DB
	cl utils.Clock
}

func NewChart(db *bun.DB, cl utils.Clock) *Chart {
	return &Chart{db: db, cl: cl}
}

func (cr *Chart) FindChartsByAuthUserId(ctx context.Context, authUserId string) ([]*domain.Chart, error) {
	var charts []*domain.Chart

	err := cr.db.NewSelect().
		Model(&charts).
		Column("label", "year", "month").
		ColumnExpr("SUM(data) AS data").
		Where("auth_user_id = ?", authUserId).
		Group("label", "year", "month").
		Order("label DESC", "year DESC", "month ASC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return charts, nil
}
