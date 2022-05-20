package main

import (
	"log"
	"simple-orm/pkg/common/config"
	"simple-orm/pkg/common/db"
	"simple-orm/pkg/products"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	products.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
