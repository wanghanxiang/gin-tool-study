package v1

import (
	"fmt"
	"net/http"
	"product-mall/internal/service"
	util "product-mall/pkg/tools"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var productService service.ProductService
	//获取文件信息
	form, _ := c.MultipartForm()
	fmt.Println("c.Request.MultipartForm", form)
	files := form.File["file"]
	//检查cookie里面的信息
	claims, _ := util.ParseToken(c.GetHeader("Cookie"))
	if err := c.ShouldBind(&productService); err == nil {
		res := productService.Create(claims.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
