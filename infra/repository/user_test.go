package repository_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/testutils"
	"github.com/taimats/bhapi/utils"
)

func TestMain(m *testing.M) {
	err := testutils.DotEnv()
	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	//Arrange
	db := testutils.SetUpDBForRepository(t)
	ctx := context.Background()
	user := &domain.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
	}

	t.Cleanup(func() {
		db.Close()
	})

	cl := utils.NewTestClocker()
	ur := repository.NewUser(db, cl)

	//Act
	got, err := ur.CreateUser(ctx, user)

	//Assert
	assert.Nil(t, err)
	assert.NotEqual(t, int64(0), got)
}

func TestFindUserByAuthUserId(t *testing.T) {
	//Arrange
	db := testutils.SetUpDBForRepository(t)
	ctx := context.Background()
	t.Cleanup(func() { db.Close() })

	user := &domain.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
	}

	cl := utils.NewTestClocker()
	ur := repository.NewUser(db, cl)

	a := assert.New(t)

	//Act
	got, err := ur.FindUserByAuthUserId(ctx, user.AuthUserId)

	//Assert
	a.Nil(err)
	a.Equal(user.Name, got.Name)
	a.Equal(user.Password, got.Password)
	a.Equal(user.Email, got.Email)
}

func TestUpdateUser(t *testing.T) {
	//Arrange
	db := testutils.SetUpDBForRepository(t)
	ctx := context.Background()
	cl := utils.NewTestClocker()

	t.Cleanup(func() { db.Close() })

	user := &domain.User{
		ID:         2,
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "update",
		Email:      domain.Email("sample@example.com"),
		Password:   domain.Password(""),
		CreatedAt:  cl.Now(),
	}

	ur := repository.NewUser(db, cl)

	a := assert.New(t)

	//Act
	err := ur.UpdateUser(ctx, user)

	//Assert
	a.Nil(err)
}

func TestDeleteUser(t *testing.T) {
	//Arrange
	db := testutils.SetUpDBForRepository(t)
	ctx := context.Background()
	cl := utils.NewTestClocker()

	t.Cleanup(func() { db.Close() })

	user := &domain.User{
		ID:         3,
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "update",
		Email:      domain.Email("sample@example.com"),
		Password:   domain.Password(""),
	}

	ur := repository.NewUser(db, cl)

	a := assert.New(t)

	//Act
	err := ur.DleteUser(ctx, user)

	//Assert
	a.Nil(err)
}
