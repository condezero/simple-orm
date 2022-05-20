package products

import (
	"simple-orm/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

// swagger:model
type UpdateProductRequestBody struct {
	// Name of the product
	Name string `json:"name"`
	// Stock of the product
	Stock int32 `json:"stock"`
	// Price of the product
	Price int32 `json:"price"`
}

// UpdateProduct godoc
// @Summary Update Products.
// @Description Update a product.
// @Tags products
// @Accept json
// @Produce json
// @Param UpdateProductRequestBody body UpdateProductRequestBody true "Product"
// @Param id path int true "Product Id"
// @Success 200 {object} models.Product
// @Router /products/{id} [put]
func (h handler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateProductRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	product.Name = body.Name
	product.Stock = body.Stock
	product.Price = body.Price

	// save product
	h.DB.Save(&product)

	return c.Status(fiber.StatusOK).JSON(&product)
}
