package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type dbConfig struct {
	host     string
	user     string
	password string
	db       string
	port     string
}

func ConnectToDB() {
	dbConfig := dbConfig{
		host:     GetEnv("POSTGRES_HOST"),
		user:     GetEnv("POSTGRES_USER"),
		password: GetEnv("POSTGRES_PASSWORD"),
		db:       GetEnv("POSTGRES_DB"),
		port:     GetEnv("POSTGRES_PORT"),
	}
	d, err := gorm.Open(postgres.Open(dbConfig.getDSN()), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}

func (cfg *dbConfig) getDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.host, cfg.user, cfg.password, cfg.db, cfg.port)
}
