package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yamaga-shu/jwt-go-next/api-service/app/router"
	"github.com/yamaga-shu/jwt-go-next/api-service/ent"
)

func main() {
	// SQLiteデータベースを持つent.Clientを作成します。
	db, err := ent.Open(dialect.SQLite, "/etc/sqlite/ent.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer db.Close()
	ctx := context.Background()
	// ent スキーマのマイグレート
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// echo の初期化
	e := echo.New()
	// # router の設定
	router.Set(e)

	e.Logger.Fatal(e.Start(":8000"))
}
