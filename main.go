package main

import (
	"github.com/adnanbrq/fiber-todo/config"
	"github.com/adnanbrq/fiber-todo/database"
)

func main() {
	cfg := config.ReadConfig()

	database.Connect(cfg)
	database.Migrate()

	app := newApp()
	app.Listen("localhost:3030")
}
