package main

import (
	"GrooveGuru/db"
	"GrooveGuru/env"
	"GrooveGuru/middleware"
	"GrooveGuru/router"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	db.Migrations()
	defer db.Close()

	app := fiber.New()
	middleware.Attach(app)
	router.Start(app.Group("/api"))

	_ = app.Listen(":" + env.Port)
}
