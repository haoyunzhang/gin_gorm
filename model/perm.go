package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Perm struct {
	gorm.Model
	PermName   string `gorm:"not null;unique"`
	Permission string
}

func PermPost(perm *Perm) error {

	ret := DB.Create(perm)

	return ret.Error
}

func PermGetOne(name string) (*Perm, error) {
	var perm Perm
	ret := DB.Where("perm_name = ?", name).Find(&perm)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &perm, nil
}

func PermGetAll() ([]Perm, error) {
	var perm []Perm
	ret := DB.Find(&perm)
	fmt.Println(perm, "------")
	if ret.Error != nil {
		return nil, ret.Error
	}
	return perm, nil
}

func PermPut(perm *Perm) error {
	ret := DB.Model(perm).Updates(perm)
	return ret.Error
}

func PermDelete(ids []int) error {
	ret := DB.Where("id IN (?)", ids).Delete(Perm{})
	return ret.Error
}
