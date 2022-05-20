package main

import (
	"log"
	"simple-orm/pkg/common/config"
	"simple-orm/pkg/common/db"
	"simple-orm/pkg/products"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "simple-orm/docs"
)

// @title Simple ORM API
// @version 1.0
// @description This is a PoC SPI.
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	products.RegisterRoutes(app, h)

	app.Get("/hc", HealthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Listen(c.Port)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /hc [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
