package handler

import "github.com/labstack/echo/v4"

type HandlerInterface interface {
	// user情報の登録
	// (POST /auth/register)
	PostAuthRegister(c echo.Context) error
	// ユーザーごとにチャートデータを返す
	// (GET /charts/{AuthUserId})
	GetChartsWithAuthUserId(c echo.Context) error
	// サーバーの監視
	// (GET /health)
	GetHealth(c echo.Context) error
	// DBサーバーの監視
	// (GET /health/db)
	GetHealthDb(c echo.Context) error
	// ユーザーごとに記録を返す
	// (GET /records/{AuthUserId})
	GetRecordsWithAuthUserId(c echo.Context) error
	// 書籍の検索結果を取得
	// (GET /search)
	GetSearch(c echo.Context) error
	// ユーザーごとに本棚を複数削除
	// (DELETE /shelf/{AuthUserId})
	DeleteShelfWithAuthUserId(c echo.Context) error
	// ユーザーごとに本棚を取得
	// (GET /shelf/{AuthUserId})
	GetShelfWithAuthUserId(c echo.Context) error
	// ユーザーごとに本を本棚に1冊ずつ作成
	// (POST /shelf/{authUserId})
	PostShelfAuthUserId(ctx echo.Context) error
	// ユーザーごとに本棚を1冊ずつ更新
	// (PUT /shelf/{AuthUserId})
	PutShelfWithAuthUserId(c echo.Context) error
	// ユーザーを削除
	// (DELETE /users/{AuthUserId})
	DeleteUsersWithAuthUserId(c echo.Context) error
	// ユーザー情報を返す
	// (GET /users/{AuthUserId})
	GetUsersWithAuthUserId(c echo.Context) error
	// ユーザー情報を更新
	// (PUT /users)
	PutUsers(c echo.Context) error
}

// Book defines model for Book.
type Book struct {
	// Id 本の識別子
	Id string `json:"id,omitempty"`

	// Isbn10 本のisbn10
	Isbn10 string `json:"isbn10,omitempty"`

	// ImageURL 本の画像
	ImageURL string `json:"imageURL,omitempty"`

	// Title 本の書名
	Title string `json:"title,omitempty"`

	// Author 本の著者
	Author string `json:"author,omitempty"`

	// Page 本のページ数
	Page string `json:"page,omitempty"`

	// Price 本の価格
	Price string `json:"price,omitempty"`

	// BookStatus 本の状態
	BookStatus string `json:"bookStatus,omitempty" validate:"required"`

	// ユーザーの識別子
	AuthUserId string `json:"authUserId,omitempty"`

	// CreatedAt 本の作成日時
	CreatedAt string `json:"createdAt,omitempty"`

	// UpdatedAt 本の更新日時
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Chart defines model for Chart.
type Chart struct {
	// Id 図表の識別子
	Id string `json:"id,omitempty"`

	// Label チャートを分類する識別子
	Label string `json:"label,omitempty"`

	// Year 各データの年
	Year string `json:"year,omitempty"`

	// Month 各データの月
	Month string `json:"month,omitempty"`

	// Data 各データ内容
	Data string `json:"data,omitempty"`
}

// Error defines model for Error.
type Error struct {
	// Message エラーメッセージ
	Message string `json:"message,omitempty"`
}

// Record defines model for Record.
type Record struct {
	// Costs 購入額の総計
	Costs string `json:"costs,omitempty"`

	// CostsRead 購入額のうち読了分
	CostsRead string `json:"costsRead,omitempty"`

	// Volumes 購入冊数の総計
	Volumes string `json:"volumes,omitempty"`

	// VolumesRead 購入冊数のうち読了分
	VolumesRead string `json:"volumesRead,omitempty"`

	// Pages 購入ページ数の総計
	Pages string `json:"pages,omitempty"`

	// PagesRead 購入ページ数のうち読了分
	PagesRead string `json:"pagesRead,omitempty"`
}

// User defines model for User.
type User struct {
	// バックユーザーの識別子
	Id string `json:"id,omitempty"`

	// AuthUserId フロントユーザーの識別子
	AuthUserId string `json:"authUserId,omitempty" validate:"required"`

	// Name ユーザー名
	Name string `json:"name,omitempty"`

	// Email ユーザーemail
	Email string `json:"email,omitempty" validate:"required,email"`

	// Password パスワード（あれば）
	Password string `json:"password,omitempty"`

	// CreatedAt ユーザーの作成日時
	CreatedAt string `json:"createdAt,omitempty"`

	// UpdatedAt ユーザーの更新日時
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type RegisterInfo struct {
	// AuthUserId ユーザーの識別子
	AuthUserId string `json:"authUserId,omitempty" validate:"required"`

	// Email ユーザーemail
	Email string `json:"email,omitempty" validate:"required,email"`

	// Name ユーザー名
	Name string `json:"name,omitempty"`

	// Password パスワード（あれば）
	Password string `json:"password,omitempty" validate:"gte=8,lte=20"`
}
