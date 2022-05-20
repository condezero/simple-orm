package products

import (
	"simple-orm/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

// DeleteProduct godoc
// @Summary Remove Products.
// @Description Remove a product from db.
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Success 200
// @Router /products/{id} [delete]
func (h handler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// delete product from db
	h.DB.Delete(&product)

	return c.SendStatus(fiber.StatusOK)
}
