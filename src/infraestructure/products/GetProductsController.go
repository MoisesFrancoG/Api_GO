package products

import (
	"demo/src/application/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProductsController struct {
	useCase_gp *product.GetProducts
}

func NewGetProductsController(useCase_gp *product.GetProducts) *GetProductsController {
	return &GetProductsController{useCase_gp: useCase_gp}
}

func (gpc *GetProductsController) Execute(c *gin.Context) {
	products, err := gpc.useCase_gp.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
