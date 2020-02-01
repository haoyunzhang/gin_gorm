package model

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);unique;index"`
	Perms     []Perm `gorm:"many2many:role_perm;"`
	PermRefer int
}

func RolePost1(c *gin.Context) {
	var perms []Perm
	DB.Find(&perms)
	role := Role{
		Name:  "角色4",
		Perms: perms,
	}

	ret := DB.Create(&role)
	if ret.Error != nil {
		c.JSON(500, gin.H{"error": ret.Error.Error()})
		return
	}
	c.JSON(201, gin.H{
		"message": "ok",
		"err":     "",
	})
}

func RolePost(role Role) error {
	var perms []Perm
	var ids []uint
	for _, v := range role.Perms {
		ids = append(ids, v.ID)
	}
	ret := DB.Where("id IN (?)", ids).Find(&perms)
	if ret.Error != nil {
		return ret.Error
	}
	role.Perms = perms
	ret = DB.Create(&role)
	return ret.Error
	if ret.Error != nil {
		return ret.Error
	}
	ass := DB.Model(&role).Association("Perms").Append(&perms)
	return ass.Error
}

//func RolePut(role Role) error {
//	var perms []Perm
//	var uRole Role
//	var ids []uint
//	for _, v := range role.Perms {
//		ids = append(ids, v.ID)
//	}
//	ret := DB.Where("id = ?", role.ID).Find(&uRole)
//	if ret.Error != nil {
//		return ret.Error
//	}
//	ret = DB.Where("id IN (?)", ids).Find(&perms)
//	if ret.Error != nil {
//		return ret.Error
//	}
//	role.Perms = nil
//	// 第一个uRole必须是查出来的
//	ret = DB.Model(&uRole).Updates(role)
//	if ret.Error != nil {
//		return ret.Error
//	}
//	// 第一个uRole必须是查出来的
//	fmt.Println(perms, "-----")
//	DB.Model(&uRole).Association("Perms").Clear()
//	ret1 := DB.Model(&uRole).Association("Perms").Replace(perms)
//	return ret1.Error
//}
func RolePut(role Role) error {
	var perms []Perm
	var uRole Role
	var ids []uint
	for _, v := range role.Perms {
		ids = append(ids, v.ID)
	}
	ret := DB.Where("id IN (?)", ids).Find(&perms)
	if ret.Error != nil {
		return ret.Error
	}
	ret = DB.Where("id = ?", role.ID).Find(&uRole)
	if ret.Error != nil {
		return ret.Error
	}
	role.Perms = nil
	// 更新的第一个role需要查出来
	ret = DB.Model(&uRole).Updates(role)
	//return ret.Error
	if ret.Error != nil {
		return ret.Error
	}

	// 这里可以添加上role对perm的关联
	DB.Model(&uRole).Association("Perms").Replace(perms)
	return nil
}

//func RolePut(role Role) error {
//	var perms []Perm
//	var uRole Role
//	var ids []uint
//	for _, v := range role.Perms {
//		ids = append(ids, v.ID)
//	}
//	ret := DB.Where("id IN (?)", ids).Find(&perms)
//	if ret.Error != nil {
//		return ret.Error
//	}
//	ret = DB.Where("id = ?", role.ID).Find(&uRole)
//	if ret.Error != nil {
//		return ret.Error
//	}
//	role.Perms = nil
//	// 更新的第一个role不是查出来的，而下面的更新关联的地方是查出来的才行，这是什么操作
//	//DB.Model(&uRole).Association("Perms").Replace(&perms)
//	ret = DB.Model(&role).Updates(role)
//	if ret.Error != nil {
//		return ret.Error
//	}
//
//	// 这里可以清除掉role对perm的关联
//	DB.Model(&role).Association("Perms").Replace(&perms)
//	// 这里可以添加上role对perm的关联
//	DB.Model(&uRole).Association("Perms").Replace(perms)
//	return nil
//}

//func RolePut(c *gin.Context) {
//	var role Role
//	var perms []Perm
//	DB.Where("name = ?", "角色1").Find(&role)
//	DB.Where("id IN (?)", []int{1,4}).Find(&perms)
//	role.Perms = perms
//	role.PermRefer = 2
//	//DB.Model(&role).Association("Perms").Clear()
//
//	ret := DB.Save(&role)
//	if ret.Error != nil {
//		c.JSON(500, gin.H{"error": ret.Error.Error()})
//		return
//	}
//	ret1 := DB.Model(&role).Association("Perms").Replace(&perms)
//	if ret1.Error != nil {
//		c.JSON(500, gin.H{"error": ret1.Error.Error()})
//		return
//	}
//	c.JSON(200, gin.H{
//		"data": "ok",
//		"err": "",
//	})
//	//DB.Model(&role).Update("perms", perms)
//}

func RoleGetOne(name string)(*Role, error) {
	//var roles Role
	//var perms []Perm
	//DB.Where("name = ?", name).Find(&roles)
	//// 这样写可以获取到该role下面关联的perms
	//ret := DB.Model(&roles).Related(&perms, "Perms")
	//if ret.Error != nil {
	//	return nil, ret.Error
	//}
	//fmt.Println(perms)
	//roles.Perms = perms
	// 下面是简单的写法
	var role Role
	ret := DB.Preload("Perms").Where("name = ?", name).Find(&role)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &role, nil
}

func RoleGetAll() ([]Role, error){

	var role []Role
	ret := DB.Preload("Perms").Find(&role)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return role, nil
}

func RoleDelete(ids []int) error {
	// 需要加上事务，否则可能出现问题
	ret := DB.Where("id IN (?)", ids).Delete(Role{})
	DB.Exec("delete from role_perm where role_id in (?)", ids)

	return ret.Error
}
