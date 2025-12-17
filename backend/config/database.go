package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host string
	Port int
	User string
	DBName string
	Password string
	SSLMode string
}

var DB *gorm.DB

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host: "localhost",
		Port: 5432,
		User: "postgres",
		Password:"mypassword" ,
		DBName: "shortenerURL_db",
		SSLMode:"disable",
	}
	return &dbConfig
}

func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
	dbConfig.Host,
	dbConfig.User,
	dbConfig.Password,
	dbConfig.DBName,
	dbConfig.Port,
	dbConfig.SSLMode)
}

func ConnectDB(){
	dbConfig := BuildDBConfig()
	dsn := DBUrl(dbConfig)

	var err error

	DB,err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	

	if err != nil {
		panic("Failed connection to the PostgreSQL Database")
	}

	fmt.Println("🥳 Database connected!")

}