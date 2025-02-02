package products

import (
	"demo/src/application/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCase_dp *product.DeleteProduct
}

func NewDeleteProductController(useCase_dp *product.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{useCase_dp: useCase_dp}
}

func (dpc *DeleteProductController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = dpc.useCase_dp.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
