package repository_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/testutils"
	"github.com/taimats/bhapi/utils"
)

func TestFindChartsByAuthUserId(t *testing.T) {
	//Arrange
	db := testutils.SetUpDBForRepository(t)
	t.Cleanup(func() { db.Close() })

	authUserId := "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058"

	cl := utils.NewTestClocker()
	cr := repository.NewChart(db, cl)

	ctx := context.Background()
	a := assert.New(t)

	//Act
	got, err := cr.FindChartsByAuthUserId(ctx, authUserId)
	log.Println(got)

	//Assert
	a.Nil(err)
	a.Len(got, 3)
}
