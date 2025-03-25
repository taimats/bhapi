package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/testutils"
	"github.com/uptrace/bun"
)

func main() {
	envPath := "C:/Users/beo03/bhapi/.env"
	relPath, err := testutils.NewRelativePath(envPath)
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load(relPath)
	if err != nil {
		log.Fatalf("envファイルの読み込みに失敗:%v", err)
	}

	conf := infra.NewDBConfig()
	db, err := infra.NewDatabaseConnection(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models := []interface{}{
		(*domain.User)(nil),
		(*domain.Book)(nil),
		(*domain.Chart)(nil),
	}

	var data []byte
	data = append(data, modelsToByte(db, models)...)

	if err = os.WriteFile("./infra/gen/schema.sql", data, 0777); err == nil {
		log.Println("DBスキーマファイルを生成")
	}
}

func modelsToByte(db *bun.DB, models []interface{}) []byte {
	var data []byte

	for _, model := range models {
		query := db.NewCreateTable().Model(model).WithForeignKeys()

		rawQuery, err := query.AppendQuery(db.Formatter(), nil)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, rawQuery...)
		data = append(data, ";\n"...)
	}

	return data
}
