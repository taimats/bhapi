package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/utils"
	"github.com/uptrace/bun"
)

type User struct {
	db *bun.DB
	cl utils.Clock
}

func NewUser(db *bun.DB, cl utils.Clock) *User {
	return &User{db: db, cl: cl}
}

// ユーザー情報の取得
func (ur *User) FindUserByAuthUserId(ctx context.Context, authUserId string) (*domain.User, error) {
	user := new(domain.User)

	err := ur.db.NewSelect().Model(user).Where("auth_user_id = ?", authUserId).Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.NewErrChains(utils.ErrNotFound, err)
		}
	}

	return user, nil
}

func (ur *User) CreateUser(ctx context.Context, user *domain.User) (userId int64, err error) {
	user.CreatedAt = ur.cl.Now()
	user.UpdatedAt = ur.cl.Now()

	tx, err := ur.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, fmt.Errorf("トランザクションの生成に失敗:%w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println(err)
		}
	}()

	err = tx.NewInsert().Model(user).Returning("id").Scan(ctx, &userId)
	if err != nil {
		return 0, fmt.Errorf("userの登録に失敗:%w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("コミット失敗:%w", err)
	}

	return userId, nil
}

func (ur *User) UpdateUser(ctx context.Context, user *domain.User) error {
	user.UpdatedAt = ur.cl.Now()

	tx, err := ur.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("トランザクションの生成に失敗:%w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println(err)
		}
	}()

	_, err = tx.NewUpdate().Model(user).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("コミット失敗:%w", err)
	}

	return nil
}

func (ur *User) DleteUser(ctx context.Context, user *domain.User) error {
	tx, err := ur.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("トランザクションの生成に失敗:%w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println(err)
		}
	}()

	_, err = tx.NewDelete().Model(user).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("コミット失敗:%w", err)
	}

	return nil
}
