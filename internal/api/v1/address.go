package v1

import (
	"product-mall/internal/service"
	util "product-mall/internal/tools"

	"github.com/gin-gonic/gin"
)

//新增收货地址
func CreateAddress(c *gin.Context) {
	service := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Cookie"))

	if err := c.BindJSON(&service); err == nil {
		res := service.Create(claim.ID)
		c.JSON(200, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(400, ErrorResponse(err))
	}
}

//展示收货地址
func ShowAddresses(c *gin.Context) {
	service := service.AddressService{}
	util.LogrusObj.WithContext(c).Infof("ShowAddresses param %s", c.Param("id"))
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

//修改收货地址
func UpdateAddress(c *gin.Context) {
	service := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Cookie"))
	if err := c.BindJSON(&service); err == nil {
		res := service.Update(claim.ID, c.Param("id"))
		c.JSON(200, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(400, ErrorResponse(err))
	}
}

//删除收获地址
func DeleteAddress(c *gin.Context) {
	service := service.AddressService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}
