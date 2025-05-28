package domain

import (
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type Password string
type Email string
type BookStatus string
type ChartLabel string

const (
	Bought  = BookStatus("bought")
	Reading = BookStatus("reading")
	Read    = BookStatus("read")
)

const (
	ChartPrice   = ChartLabel("購入額")
	ChartVolumes = ChartLabel("購入冊数")
	ChartPages   = ChartLabel("購入ページ数")
)

func (p Password) String() string {
	return "xxxxxxxxx"
}
func (p Password) GoString() string {
	return "xxxxxxxxx"
}
func (e Email) String() string {
	return "xxxx@xxxx"
}
func (e Email) GoString() string {
	return "xxxx@xxxx"
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID         int64     `bun:",pk,autoincrement"`
	AuthUserId string    `bun:"auth_user_id,nullzero,notnull,unique"`
	Name       string    `bun:"name"`
	Email      Email     `bun:"email,notnull,unique"`
	Password   Password  `bun:"password"`
	CreatedAt  time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt  time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Book struct {
	bun.BaseModel `bun:"table:books,alias:b"`

	ID         int64      `bun:",pk,autoincrement" json:"id,omitempty"`
	ISBN10     string     `bun:"isbn_10" json:"isbn10,omitempty"`
	ImageURL   string     `bun:"image_url" json:"imageURL,omitempty"`
	Title      string     `bun:"title" json:"title,omitempty"`
	Author     string     `bun:"author" json:"author,omitempty"`
	Page       int        `bun:"page,type:integer" json:"page,omitempty"`
	Price      int        `bun:"price,type:integer" json:"price,omitempty"`
	BookStatus BookStatus `bun:"book_status,nullzero,notnull" json:"bookStatus,omitempty"`
	AuthUserId string     `bun:"auth_user_id,nullzero,notnull" json:"authUserId,omitempty"`
	CreatedAt  time.Time  `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt  time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updatedAt,omitempty"`
}

type Record struct {
	Costs       int `json:"costs,omitempty"`
	CostsRead   int `json:"costsRead,omitempty"`
	Volumes     int `json:"volumes,omitempty"`
	VolumesRead int `json:"volumesRead,omitempty"`
	Pages       int `json:"pages,omitempty"`
	PagesRead   int `json:"pagesRead,omitempty"`
}

type Chart struct {
	bun.BaseModel `bun:"table:charts,alias:c"`

	ID         int64      `bun:",pk,autoincrement" json:"id,omitempty"`
	Label      ChartLabel `bun:"label,nullzero" json:"label,omitempty"`
	Year       int        `bun:"year,nullzero,type:integer" json:"year,omitempty"`
	Month      int        `bun:"month,nullzero,type:integer" json:"month,omitempty"`
	Data       int        `bun:"data,nullzero,type:integer" json:"data,omitempty"`
	AuthUserId string     `bun:"auth_user_id,nullzero,notnull" json:"authUserId,omitempty"`
	BookId     int64      `bun:"book_id,nullzero,notnull" json:"bookId,omitempty"`
	CreatedAt  time.Time  `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt  time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updatedAt,omitempty"`
}

// パスワードをハッシュ化する。内部でbcryptパッケージを使用しており、Costはデフォルトの10で固定。
// errorはbcryptパッケージのエラーに一致する（この関数自体の独自エラーはなし）。
func (u *User) HashedPassword(pw Password) (Password, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return Password(hp), nil
}

// 入力されたパスワード(raw)がハッシュ化されたパスワードに一致するかを検証。
func (u *User) ValidatePassword(raw Password) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(raw)); err != nil {
		return false
	}
	return true
}

func NewChartsFromBook(book *Book) []*Chart {
	year := book.CreatedAt.Year()
	month := int(book.CreatedAt.Month())

	chartPrice := &Chart{
		Label:      ChartPrice,
		Year:       year,
		Month:      month,
		Data:       book.Price,
		AuthUserId: book.AuthUserId,
	}

	chartVolume := &Chart{
		Label:      ChartVolumes,
		Year:       year,
		Month:      month,
		Data:       1,
		AuthUserId: book.AuthUserId,
	}

	chartPage := &Chart{
		Label:      ChartPages,
		Year:       year,
		Month:      month,
		Data:       book.Page,
		AuthUserId: book.AuthUserId,
	}

	charts := []*Chart{chartPrice, chartVolume, chartPage}

	return charts
}

func UpdateChartsFromBook(book *Book) []*Chart {
	year := book.CreatedAt.Year()
	month := int(book.CreatedAt.Month())

	chartPrice := &Chart{
		Label:      ChartPrice,
		Year:       year,
		Month:      month,
		Data:       book.Price,
		AuthUserId: book.AuthUserId,
		BookId:     book.ID,
		CreatedAt:  book.CreatedAt,
	}

	chartVolume := &Chart{
		Label:      ChartVolumes,
		Year:       year,
		Month:      month,
		Data:       1,
		AuthUserId: book.AuthUserId,
		BookId:     book.ID,
		CreatedAt:  book.CreatedAt,
	}

	chartPage := &Chart{
		Label:      ChartPages,
		Year:       year,
		Month:      month,
		Data:       book.Page,
		AuthUserId: book.AuthUserId,
		BookId:     book.ID,
		CreatedAt:  book.CreatedAt,
	}

	charts := []*Chart{chartPrice, chartVolume, chartPage}

	return charts
}

func NewRecordFromBooks(books []*Book) *Record {
	record := new(Record)

	record.Volumes = len(books)

	for _, b := range books {
		record.Costs += b.Price
		record.Pages += b.Page
		if b.BookStatus == Read {
			record.CostsRead += b.Price
			record.VolumesRead += 1
			record.PagesRead += b.Page
		}
	}

	return record
}

//****** データベースへの保存を断念（念のためタグは保存）******
// type Record struct {
// 	bun.BaseModel `bun:"table:records,alias:r"`

// 	ID          int64      `bun:",pk,autoincrement"`
// 	Costs       int        `bun:"costs,notnull"`
// 	CostsRead   int        `bun:"costs_read,notnull"`
// 	Volumes     int        `bun:"volumes,notnull"`
// 	VolumesRead int        `bun:"volumes_read,notnull"`
// 	Pages       int        `bun:"pages,notnull"`
// 	PagesRead   int        `bun:"pages_read,notnull"`
// 	AuthUserId  string     `bun:"auth_user_id,nullzero,notnull,unique"`
// 	BookStatus  BookStatus `bun:"book_status,nullzero,notnull,unique"`
// 	CreatedAt   time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
// 	UpdatedAt   time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
// }
