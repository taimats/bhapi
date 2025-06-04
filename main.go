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

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"github.com/taimats/bhapi/controller"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/presenter/handler"
	"github.com/taimats/bhapi/presenter/middleware"
	"github.com/taimats/bhapi/utils"
)

func main() {
	//環境変数の設定
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(".envファイルの読み込みに失敗:%s", err)
	}

	//データベースの接続設定
	dsn := infra.NewPostgresDsn()
	db, err := infra.NewBunDB(dsn)
	if err != nil {
		log.Fatalf("データベースの接続に失敗:%s", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

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
	hc := controller.NewHealthDB(db)

	//hanlderの生成
	h := handler.NewHandler(uc, cc, rc, sc, sbc, hc)

	//echoの生成
	e := middleware.SetAll(echo.New())
	handler.RegisterHandlersWithBaseURL(e, h)

	//サーバーの起動
	address := fmt.Sprintf("%v:%v", os.Getenv("BACK_API_HOST"), os.Getenv("BACK_API_PORT"))
	go func() {
		if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatalf("サーバーの起動に失敗:%w", err)
		}
	}()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	//サーバーのシャットダウンの処理
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("サーバーのシャットダウンに失敗:%w", err)
	}
	log.Println("サーバーが正常にシャットダウンしました")
}
