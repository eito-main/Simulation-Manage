package config

// DBConfig はデータベース接続設定を管理する構造体
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

// GetDBConfig はデータベース設定を取得する
func GetDBConfig() DBConfig {
	return DBConfig{
		User:     "postgres",
		Password: "postgres",
		Host:     "postgres_container",
		Port:     "5432",
		Name:     "simulatio-manage",
		SSLMode:  "disable",
	}
}
