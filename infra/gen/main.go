package main

import (
	"log"
	"os"

	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra"
	"github.com/taimats/bhapi/testutils"
	"github.com/uptrace/bun"
)

func main() {
	err := testutils.DotEnv()
	if err != nil {
		log.Fatalf(".envファイルの取得に失敗:%s", err)
	}

	dsn := infra.NewPostgresDsn()
	bundb, err := infra.NewBunDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer bundb.Close()

	models := []interface{}{
		(*domain.User)(nil),
		(*domain.Book)(nil),
		(*domain.Chart)(nil),
	}

	var data []byte
	data = append(data, modelsToByte(bundb, models)...)

	if err = os.WriteFile("./infra/gen/schema.sql", data, 0777); err == nil {
		log.Println("DBスキーマファイルを生成")
	}
}

func modelsToByte(bundb *bun.DB, models []interface{}) []byte {
	var data []byte

	for _, model := range models {
		query := bundb.NewCreateTable().Model(model).WithForeignKeys()

		rawQuery, err := query.AppendQuery(bundb.Formatter(), nil)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, rawQuery...)
		data = append(data, ";\n"...)
	}

	return data
}
