package model

import (
	"database/sql"
	"fmt"
)

// User はユーザ情報を保持する構造体
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

// GetUserByID は指定された ID のユーザ情報を取得
func GetUserByID(id int) (*User, error) {
	query := "SELECT id, name, age, gender FROM users WHERE id = $1"
	row := DB.QueryRow(query, id) // `model.DB` を直接参照

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
