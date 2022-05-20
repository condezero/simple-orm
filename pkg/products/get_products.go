package products

import (
	"simple-orm/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

// GetProduct godoc
// @Summary Get all Products.
// @Description Get all products from db.
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func (h handler) GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&products)
}
