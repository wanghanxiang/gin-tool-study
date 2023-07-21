package serializer

import (
	"product-mall/model"
)

type Product struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	CategoryID       uint   `json:"category_id"`
	Title            string `json:"title"`
	Info             string `json:"info"`
	ImgPath          string `json:"img_path"`
	Price            string `json:"price"`
	DiscoutPrice     string `json:"discount_price"`
	View             uint64 `json:"view"`
	CreatedAt        int64  `json:"created_at"`
	Num              int    `json:"num"`
	OnSale           bool   `json:"on_sale"`
	CreateUserID     int    `json:"create_user_id"`
	CreateUserName   string `json:"create_user_name"`
	CreateUserAvatar string `json:"create_user_avatar"`
}

// 序列化商品
func BuildProduct(item model.Product) Product {
	return Product{
		ID:               item.ID,
		Name:             item.Name,
		CategoryID:       item.CategoryID,
		Title:            item.Title,
		Info:             item.Info,
		ImgPath:          item.ImgPath,
		Price:            item.Price,
		DiscoutPrice:     item.DiscountPrice,
		View:             item.GetView(),
		Num:              item.Num,
		OnSale:           item.OnSale,
		CreatedAt:        item.CreatedAt.Unix(),
		CreateUserID:     item.CreateUserID,
		CreateUserName:   item.CreateUserName,
		CreateUserAvatar: item.CreateUserAvatar,
	}
}

//序列化商品列表
func BuildProducts(items []model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
