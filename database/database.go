package database

import (
	"fmt"

	"github.com/adnanbrq/fiber-todo/config"
	"github.com/adnanbrq/fiber-todo/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB instance
var DB *gorm.DB = nil

// Connect will try to establish a connection to the given postgres database
// if a connection is already active then it will do nothing
func Connect(cfg config.Config) {
	if DB != nil {
		return
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable host=%s",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbName,
		cfg.DbPort,
		cfg.DbHost,
	)

	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing: false,
	})

	if err != nil {
		panic(err)
	}

	DB = con
}

// Migrate runs gotm AutoMigration
func Migrate() {
	DB.AutoMigrate(&model.Todo{})
}
