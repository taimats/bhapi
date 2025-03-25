package controller_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/testutils"
	"github.com/taimats/bhapi/utils"
)

func TestRegisterUser(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	container, Terminate, err := testutils.SetUpDBContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { Terminate(container.Container) })

	err = testutils.MigrateUp(container.DSN, MigrationsPath)
	if err != nil {
		t.Fatalf("マイグレーションに失敗:%s", err)
	}

	db := testutils.SetUpDBForController(t, container.DSN)
	cl := utils.NewTestClocker()
	ur := repository.NewUser(db, cl)
	sut := controller.NewUser(ur)

	user := &domain.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
	}

	a := assert.New(t)

	//Act ***************
	err = sut.RegisterUser(ctx, user)

	//Assert ***************
	a.Nil(err)
}

func TestGetUser(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	container, Terminate, err := testutils.SetUpDBContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { Terminate(container.Container) })

	err = testutils.MigrateUp(container.DSN, MigrationsPath)
	if err != nil {
		t.Fatalf("マイグレーションに失敗:%s", err)
	}

	db := testutils.SetUpDBForController(t, container.DSN)
	cl := utils.NewTestClocker()
	ur := repository.NewUser(db, cl)
	sut := controller.NewUser(ur)

	user := &domain.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
	}

	testutils.InsertTestData(t, db, ctx, user)

	a := assert.New(t)

	//Act ***************
	got, err := sut.GetUser(ctx, user.AuthUserId)

	//Assert ***************
	a.Nil(err)
	a.Equal(user.AuthUserId, got.AuthUserId)
	a.Equal(user.Name, got.Name)
	a.Equal(user.Email, got.Email)
	a.Equal(user.Password, got.Password)
}

func TestUpdateUser(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	container, Terminate, err := testutils.SetUpDBContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { Terminate(container.Container) })

	err = testutils.MigrateUp(container.DSN, MigrationsPath)
	if err != nil {
		t.Fatalf("マイグレーションに失敗:%s", err)
	}

	db := testutils.SetUpDBForController(t, container.DSN)
	cl := utils.NewTestClocker()
	ur := repository.NewUser(db, cl)
	sut := controller.NewUser(ur)

	user := &domain.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
	}
	testutils.InsertTestData(t, db, ctx, user)

	updatedUser := &domain.User{
		ID:         int64(1),
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "update",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
		CreatedAt:  cl.Now(),
	}

	a := assert.New(t)

	//Act ***************
	err = sut.UpdateUser(ctx, updatedUser)

	//Assert ***************
	a.Nil(err)
}

func TestDeleteUser(t *testing.T) {
	//Arrange ***************
	ctx := context.Background()
	container, Terminate, err := testutils.SetUpDBContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { Terminate(container.Container) })

	err = testutils.MigrateUp(container.DSN, MigrationsPath)
	if err != nil {
		t.Fatalf("マイグレーションに失敗:%s", err)
	}

	db := testutils.SetUpDBForController(t, container.DSN)
	cl := utils.NewTestClocker()
	ur := repository.NewUser(db, cl)
	sut := controller.NewUser(ur)

	user := &domain.User{
		AuthUserId: "c0cc3f0c-9a02-45ba-9de7-7d7276bb6058",
		Name:       "",
		Email:      domain.Email("example@example.com"),
		Password:   domain.Password(""),
	}
	testutils.InsertTestData(t, db, ctx, user)

	a := assert.New(t)

	//Act ***************
	err = sut.DeleteUser(ctx, user.AuthUserId)

	//Assert ***************
	a.Nil(err)
}
