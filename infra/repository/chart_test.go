package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/testutils"
)

func TestFindChartsByAuthUserId(t *testing.T) {
	//Arrange
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer bundb.Close()

	charts := []*domain.Chart{
		{
			ID:         int64(1),
			Label:      domain.ChartPrice,
			Year:       2025,
			Month:      2,
			Data:       980,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(2),
			Label:      domain.ChartVolumes,
			Year:       2025,
			Month:      2,
			Data:       1,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(3),
			Label:      domain.ChartPages,
			Year:       2025,
			Month:      2,
			Data:       247,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(4),
			Label:      domain.ChartPrice,
			Year:       2025,
			Month:      2,
			Data:       980,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(5),
			Label:      domain.ChartVolumes,
			Year:       2025,
			Month:      2,
			Data:       1,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
		{
			ID:         int64(6),
			Label:      domain.ChartPages,
			Year:       2025,
			Month:      2,
			Data:       247,
			AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
			BookId:     int64(1),
			CreatedAt:  cl.Now(),
			UpdatedAt:  cl.Now(),
		},
	}
	testutils.InsertTestData(ctx, t, bundb, charts...)

	want := []*domain.Chart{
		{
			Label: domain.ChartPrice,
			Year:  2025,
			Month: 2,
			Data:  1960,
		},
		{
			Label: domain.ChartVolumes,
			Year:  2025,
			Month: 2,
			Data:  2,
		},
		{
			Label: domain.ChartPages,
			Year:  2025,
			Month: 2,
			Data:  494,
		},
	}
	authUserId := "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058"
	sut := repository.NewChart(bundb, cl)

	a := assert.New(t)

	//Act
	got, err := sut.FindChartsByAuthUserId(ctx, authUserId)

	//Assert
	a.Nil(err)
	a.Equal(want, got)
}
