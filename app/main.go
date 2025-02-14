package main

import (
	"Simulation-Manage/config"
	"Simulation-Manage/server"
	"log"
)

func main() {
	// データベース接続の初期化
	config.InitDB()

	// サーバの起動
	log.Println("サーバを起動中...")
	server.Run()
}