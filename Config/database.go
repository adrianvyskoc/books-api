package Config

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host			string
	Port			int
	User			string
	DBName		string
	Password	string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:			"localhost",
		Port:			5432,
		User:			"adrianvyskoc",
		Password:	"12345",
		DBName:		"books",
	}
	return &dbConfig
}

func DbUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}