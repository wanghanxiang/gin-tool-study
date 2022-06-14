package service

import (
	"gin-tool-study/model"
	"gin-tool-study/pkg/e"
	"gin-tool-study/serializer"
	"mime/multipart"

	logging "github.com/sirupsen/logrus"
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
func (service *ProductService) Create(id uint, files []*multipart.FileHeader) serializer.Response {
	code := e.SUCCESS
	//获取用户信息
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
		}
	}
	tmp, _ := files[0].Open()
	status, info := Upload2QiNiu(tmp, files[0].Size)
	if status != 200 {
		return serializer.Response{
			Status: status,
			Data:   e.GetMsg(status),
			Error:  info,
		}
	}
	//存储product
	product := model.Product{
		Name:             service.Name,
		CategoryID:       uint(service.CategoryID),
		Title:            service.Title,
		Info:             service.Info,
		ImgPath:          info,
		Price:            service.Price,
		DiscountPrice:    service.DiscountPrice,
		Num:              service.Num,
		OnSale:           true,
		CreateUserID:     int(id),
		CreateUserName:   user.UserName,
		CreateUserAvatar: user.Avatar,
	}
	err := model.DB.Create(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//解析文件信息--其他文件存数据库-商品id和文件地址保存在一起
	for _, file := range files {
		tmp, _ := file.Open()
		status, info := Upload2QiNiu(tmp, file.Size)
		if status != 200 {
			return serializer.Response{
				Status: status,
				Data:   e.GetMsg(status),
				Error:  info,
			}
		}
		productImg := model.ProductImg{
			ProductID: product.ID,
			ImgPath:   info,
		}
		err = model.DB.Create(&productImg).Error
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}

}
