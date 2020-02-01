package model

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Age      int
	Name     string `gorm:"size:255;unique"` // string默认长度为255, 使用这种tag重设。
	PassWord string
	Num      int
	Roles    []Role `gorm:"many2many:user_role;"` //
}

func UserPost(user User) error {
	var roles []Role
	var ids []uint
	for _, v := range user.Roles {
		ids = append(ids, v.ID)
	}
	ret := DB.Where("id IN (?)", ids).Find(&roles)
	if ret.Error != nil {
		return ret.Error
	}
	user.Roles = roles
	user.PassWord = GenerateId(User{Name: user.Name, PassWord: user.PassWord})
	ret = DB.Create(&user)
	return ret.Error
}

func UserGetOne(name string) (*User, error) {
	var user User
	ret := DB.Preload("Roles").Where("name = ?", name).Find(&user)
	for k, val := range user.Roles {
		var perms []Perm
		DB.Model(&val).Related(&perms, "Perms")
		user.Roles[k].Perms = perms
	}
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &user, nil
}

func UserGetAll() ([]User, error) {

	var users []User
	ret := DB.Preload("Roles").Find(&users)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return users, nil
}

func UserPut(user User) error {
	var roles []Role
	var uUser User
	var ids []uint
	for _, v := range user.Roles {
		ids = append(ids, v.ID)
	}
	ret := DB.Where("id IN (?)", ids).Find(&roles)
	if ret.Error != nil {
		return ret.Error
	}
	ret = DB.Where("id = ?", user.ID).Find(&uUser)
	if ret.Error != nil {
		return ret.Error
	}
	user.Roles = nil
	// 更新密码
	if user.PassWord != "" {
		if user.Name != "" {
			user.PassWord = GenerateId(User{Name: user.Name, PassWord: user.PassWord})
		}else {
			user.PassWord = GenerateId(User{Name: uUser.Name, PassWord: user.PassWord})
		}
	}
	// 更新的第一个role需要查出来
	ret = DB.Model(&uUser).Updates(user)
	if ret.Error != nil {
		return ret.Error
	}

	// 这里可以添加上role对perm的关联
	DB.Model(&uUser).Association("Roles").Replace(roles)
	return nil
}

func UserDelete(ids []int) error {
	// 需要加上事务，否则可能出现问题
	ret := DB.Where("id IN (?)", ids).Delete(User{})
	DB.Exec("delete from user_role where user_id in (?)", ids)

	return ret.Error
}

func GenerateId(inst interface{}) string {

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err := enc.Encode(inst)
	if err != nil {
		panic(err.Error())
	}
	return fmt.Sprintf("%x", md5.Sum(buf.Bytes()))
}
