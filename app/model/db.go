package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQLドライバ
	"Simulation-Manage/config"
)

// DB はデータベース接続を保持する変数
var DB *sql.DB

// InitDB はPostgreSQLへの接続を初期化
func InitDB() {
	cfg := config.GetDBConfig() // 設定を取得

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("データベースの接続エラー: ", err)
	}

	// 接続の確認
	if err = DB.Ping(); err != nil {
		log.Fatal("データベースへの接続確認エラー: ", err)
	}

	fmt.Println("データベース接続成功")
}
