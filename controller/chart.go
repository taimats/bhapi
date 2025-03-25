package controller

import (
	"context"

	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra/repository"
)

type Chart struct {
	cr *repository.Chart
}

func NewChart(cr *repository.Chart) *Chart {
	return &Chart{cr: cr}
}

func (cc *Chart) GetCharts(ctx context.Context, authUserId string) ([]*domain.Chart, error) {
	charts, err := cc.cr.FindChartsByAuthUserId(ctx, authUserId)
	if err != nil {
		return nil, err
	}

	return charts, nil
}
