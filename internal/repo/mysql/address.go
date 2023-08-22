package mysql

import (
	"product-mall/internal/model"
	"product-mall/pkg/db"

	"gorm.io/gorm"
)

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo() *addressRepo {
	return &addressRepo{
		db: db.GetDB(),
	}
}

func (a *addressRepo) Create(address *model.Address) error {
	return a.db.Create(address).Error
}

func (a *addressRepo) GetAddressByUid(user_id interface{}) (address []model.Address, err error) {
	err = a.db.Model(model.Address{}).Where("user_id = ?", user_id).Order("created_at DESC").Find(&address).Error
	return
}

func (a *addressRepo) GetAddressById(id string) (address model.Address, err error) {
	err = db.GetDB().Where("id = ?", id).Find(&address).Error
	return
}

func (a *addressRepo) DeleteAddress(address model.Address) error {
	err := db.GetDB().Delete(&address).Error
	return err
}

// 更新对应的address
func (a *addressRepo) Updates(address *model.Address) error {
	return a.db.Model(&address).Updates(&address).Error
}
