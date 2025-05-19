package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/testutils"
	"github.com/taimats/bhapi/utils"
)

func TestCreateUser(t *testing.T) {
	//Arrange
	ctx := context.Background()
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	dbctr.Restore(ctx, t)
	cl := utils.NewTestClocker()

	user := &domain.User{
		ID:         int64(1),
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
	}
	sut := repository.NewUser(bundb, cl)
	a := assert.New(t)

	//Act
	got, err := sut.CreateUser(ctx, user)

	//Assert
	a.Equal(int64(1), got)
	a.Nil(err)
}

func TestFindUserByAuthUserId(t *testing.T) {
	//Arrange
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	cl := utils.NewTestClocker()

	user := &domain.User{
		ID:         int64(1),
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
		CreatedAt:  cl.Now(),
		UpdatedAt:  cl.Now(),
	}
	testutils.InsertTestData(ctx, t, bundb, user)
	sut := repository.NewUser(bundb, cl)

	a := assert.New(t)

	//Act
	got, err := sut.FindUserByAuthUserId(ctx, user.AuthUserId)

	//Assert
	a.Nil(err)
	a.Equal(user, got)
}

func TestUpdateUser(t *testing.T) {
	//Arrange
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	cl := utils.NewTestClocker()

	user := &domain.User{
		ID:         int64(1),
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
		CreatedAt:  cl.Now(),
		UpdatedAt:  cl.Now(),
	}
	testutils.InsertTestData(ctx, t, bundb, user)

	updatedUser := &domain.User{
		ID:         int64(1),
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "update",
		Email:      domain.Email("sample@example.com"),
		Password:   domain.Password(""),
		CreatedAt:  cl.Now(),
		UpdatedAt:  cl.Now(),
	}

	sut := repository.NewUser(bundb, cl)

	a := assert.New(t)

	//Act
	err = sut.UpdateUser(ctx, updatedUser)

	//Assert
	a.Nil(err)
}

func TestDeleteUser(t *testing.T) {
	//Arrange
	ctx := context.Background()
	dbctr.Restore(ctx, t)
	bundb, err := infra.NewBunDB(dbctr.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	cl := utils.NewTestClocker()

	user := &domain.User{
		ID:         int64(1),
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
		CreatedAt:  cl.Now(),
		UpdatedAt:  cl.Now(),
	}
	testutils.InsertTestData(ctx, t, bundb, user)

	sut := repository.NewUser(bundb, cl)

	a := assert.New(t)

	//Act
	err = sut.DleteUser(ctx, user)

	//Assert
	a.Nil(err)
}
