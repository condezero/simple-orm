package products

import (
	"simple-orm/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

// swagger:model
type AddProductRequestBody struct {
	// Name of the product
	Name string `json:"name"`
	// Stock of the product
	Stock int32 `json:"stock"`
	// Price of the product
	Price int32 `json:"price"`
}

// AddProducts godoc
// @Summary Add Products.
// @Description Adds the product to db.
// @Tags products
// @Accept json
// @Produce json
// @Param AddProductRequestBody body AddProductRequestBody true "Product"
// @Success 200 {object} models.Product
// @Router /products [post]
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
