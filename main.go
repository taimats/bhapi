package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/presenter/handler"
	"github.com/taimats/bhapi/presenter/middleware/auth"
	"github.com/taimats/bhapi/utils"
)

var (
	allowedOrigins = []string{os.Getenv("FRONT_API_BASE_URL")}
	allowedMethods = []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete}
	allowedHeaders = []string{echo.HeaderContentType, echo.HeaderAuthorization}
)

func main() {
	// /************ローカル環境での確認用****************/
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf(".envファイルの読み込みに失敗:%s", err)
	// }

	//データベースの接続設定
	dsn := infra.NewDBConfig()
	db, err := infra.NewDatabaseConnection(dsn)
	if err != nil {
		log.Fatalf("データベースの接続に失敗:%s", err)
	}
	defer db.Close()

	//repositoryインスタンスの生成
	cl := utils.NewClocker()
	cr := repository.NewChart(db, cl)
	sr := repository.NewShelf(db, cl)
	ur := repository.NewUser(db, cl)

	//controllerインスタンスの生成
	cc := controller.NewChart(cr)
	sc := controller.NewShelf(sr)
	uc := controller.NewUser(ur)
	rc := controller.NewRecord(sr)
	sbc := controller.NewSearchBooks()

	e := echo.New()

	//echoのmiddlewareの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowMethods: allowedMethods,
		AllowHeaders: allowedHeaders,
	}))
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:    middleware.DefaultSkipper,
		KeyLookup:  "header:" + echo.HeaderAuthorization,
		AuthScheme: "Bearer",
		Validator: func(key string, c echo.Context) (bool, error) {
			ok, err := auth.Authenticate(key)
			return ok, err
		},
	}))
	e.Validator = handler.NewCustomValidator(validator.New())

	//echoのhanlderの設定
	server := handler.NewServer(uc, cc, rc, sc, sbc)
	handler.RegisterHandlersWithBaseURL(e, server)

	//graceful shutdownの設定
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	//サーバーの起動
	address := fmt.Sprintf("%v:%v", os.Getenv("BACK_API_HOST"), os.Getenv("BACK_API_PORT"))
	go func() {
		if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatalf("サーバーの起動に失敗:%w", err)
		}
	}()

	//サーバーのシャットダウンの処理
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("サーバーのシャットダウンに失敗:%w", err)
	}

	log.Println("サーバーが正常にシャットダウンしました")
}
