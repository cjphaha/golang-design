package di

import "database/sql"

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

// NewConfig 创建配置
func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "./example.db",
		Port:         "8000",
	}
}

// ConnectDatabase 打开数据库连接
func ConnectDatabase(config *Config) (*sql.DB, error) {
	return sql.Open("sqlite3", config.DatabasePath)
}
