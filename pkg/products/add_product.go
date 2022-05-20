package products

import (
	"simple-orm/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type AddProductRequestBody struct {
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
	Price int32  `json:"price"`
}

func (h handler) AddProduct(c *fiber.Ctx) error {
	body := AddProductRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var product models.Product
	product.Name = body.Name
	product.Stock = body.Stock
	product.Price = body.Price

	if result := h.DB.Create(&product); result.Error != nil {
		return fiber.NewError(fiber.StatusConflict, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&product)
}
