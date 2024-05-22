package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DbConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
}

func loadEnvDb() DbConfig {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	fmt.Printf("Parsing %s\t%s\t%s\t%s\t%s\n", dbUser, dbPassword, dbName, dbPort, dbHost)
	return DbConfig{
		Host:     dbHost,
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
		Port:     dbPort,
	}
}

func ConnectDB() *gorm.DB {
	dbConfig := loadEnvDb()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Bucharest",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}

	return db
}
