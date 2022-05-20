package products

import (
	"simple-orm/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

// GetProduct godoc
// @Summary Get a Product.
// @Description Get a product from db.
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func (h handler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&product)
}
