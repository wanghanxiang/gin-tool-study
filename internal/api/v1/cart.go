package v1

import (
	"net/http"
	"product-mall/internal/service"
	util "product-mall/internal/tools"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	createCartService := service.CartService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(c.Param("id"), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ShowList(c *gin.Context) {
	showCartsService := service.CartService{}
	res := showCartsService.List(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func UpdateCart(c *gin.Context) {
	updateCartService := service.CartService{}
	if err := c.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.Update(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteCart(c *gin.Context) {
	deleteCartService := service.CartService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(c.Param("id"), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
