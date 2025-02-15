package main

import (
  "Simulation-Manage/model"
	"Simulation-Manage/server"
	"log"
)

func main() {
	// データベース接続の初期化
	model.InitDB()

	// サーバの起動
	log.Println("サーバを起動中...")
	server.Run()
}