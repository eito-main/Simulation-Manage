package model

import (
	"Simulation-Manage/config"
	"database/sql"
	"fmt"
)

// Userはユーザ情報を保持する構造体です
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

// GetUserByIDは指定されたIDのユーザ情報をデータベースから取得します
func GetUserByID(id int) (*User, error) {
	query := "SELECT id, name, age, gender FROM users WHERE id = $1"
	row := config.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ユーザが見つかりません: ID=%d", id)
		}
		return nil, fmt.Errorf("データベースエラー: %v", err)
	}

	return &user, nil
}
