package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/utils"
	"github.com/uptrace/bun"
)

// book, chartのテーブルをトランザクションで操作するshelf構造体
type Shelf struct {
	db *bun.DB
	cl utils.Clock
}

func NewShelf(db *bun.DB, cl utils.Clock) *Shelf {
	return &Shelf{db: db, cl: cl}
}

// authUserIdをもとに本棚を返す
func (sr *Shelf) FindBooksByAuthUserID(ctx context.Context, authUserId string) ([]*domain.Book, error) {
	var books []*domain.Book

	err := sr.db.NewSelect().Model(&books).Where("auth_user_id = ?", authUserId).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

// 本を新規で作成時に、book、chartsをまとめてデータベースに登録
func (sr *Shelf) CreateBookWithCharts(ctx context.Context, book *domain.Book, charts []*domain.Chart) error {
	now := sr.cl.Now()

	book.CreatedAt = now
	book.UpdatedAt = now
	for _, c := range charts {
		c.CreatedAt = now
		c.UpdatedAt = now
	}

	//トランザクション
	tx, err := sr.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("トランザクションの生成に失敗:%w", err)
	}
	defer tx.Rollback()

	//本の登録
	var bookId int64
	err = tx.NewInsert().Model(book).Returning("id").Scan(ctx, &bookId)
	if err != nil {
		return err
	}

	//チャートの登録
	for _, c := range charts {
		c.BookId = bookId
	}
	err = tx.NewInsert().Model(&charts).Scan(ctx)
	if err != nil {
		return err
	}

	//コミット処理
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("コミット失敗:%w", err)
	}

	return nil
}

// 本の更新時（チャートの更新は不要）
func (sr *Shelf) UpdateBook(ctx context.Context, book *domain.Book) error {
	book.UpdatedAt = sr.cl.Now()

	//トランザクション
	tx, err := sr.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("トランザクションの生成に失敗:%w", err)
	}
	defer tx.Rollback()

	//本の更新
	_, err = tx.NewUpdate().Model(book).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	//コミット処理
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("コミット失敗:%w", err)
	}

	return nil
}

// 本の削除時、book_idで対応するチャートも削除
func (sr *Shelf) DleteBooksWithCharts(ctx context.Context, books []*domain.Book) error {
	var bookIds []int64
	for _, b := range books {
		bookIds = append(bookIds, b.ID)
	}

	//トランザクションの開始
	tx, err := sr.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("トランザクションの生成に失敗:%w", err)
	}
	defer tx.Rollback()

	//本の削除
	err = tx.NewDelete().Model(&books).WherePK().Scan(ctx)
	if err != nil {
		return err
	}

	//bookIdをもとにチャートを削除
	//NewDelete where inでは削除に失敗するので、仕方なく遠回りな実装となっている
	var charts []*domain.Chart
	err = tx.NewSelect().Model(&charts).Column("id").Where("book_id IN (?)", bun.In(bookIds)).Scan(ctx)
	if err != nil {
		return err
	}
	err = tx.NewDelete().Model(&charts).WherePK().Scan(ctx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("コミット失敗:%w", err)
	}

	return nil
}
