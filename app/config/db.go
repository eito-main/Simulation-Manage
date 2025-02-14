package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQLドライバ
)

var DB *sql.DB

// InitDBはPostgreSQLへの接続を初期化します
func InitDB() {
	var err error
	connStr := "postgres://postgres:postgres@postgres_container:5432/simulatio-manage?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("データベースの接続エラー: ", err)
	}

	// 接続の確認
	err = DB.Ping()
	if err != nil {
		log.Fatal("データベースへの接続確認エラー: ", err)
	}

	fmt.Println("データベース接続成功")
}
