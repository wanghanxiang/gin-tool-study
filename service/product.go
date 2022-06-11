package service

import (
	"gin-tool-study/pkg/e"
	"gin-tool-study/serializer"
)

/*
创建商品信息
**/

type ProductService struct {
	ID            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale        bool   `form:"on_sale" json:"on_sale"`
	Num           int    `form:"num" json:"num"`
	PageNum       int    `form:"pageNum"`
	PageSize      int    `form:"pageSize"`
}

//创建商品
func (service *ProductService) Create() serializer.Response {
	code := e.SUCCESS

	//获取用户信息

	//解析文件信息

	//商品信息存库，商品详细图片存库

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
