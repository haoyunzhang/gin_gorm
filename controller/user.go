package controller

import (
	"github.com/gin-gonic/gin"
	"hfga.com.cn/ginorm/model"
	"net/http"
	"strconv"
)

func UserPost(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ret := model.UserPost(user)
	if ret != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ret.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"error": ""})
}

func UserGet(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		userGetAll(c)
		return
	}
	userGetOne(c, name)
}

func userGetOne(c *gin.Context, name string) {
	ret, err := model.UserGetOne(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ret, "error": ""})
}

func userGetAll(c *gin.Context) {
	ret, err := model.UserGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ret, "error": ""})
}

func UserPut(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := model.UserPut(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"error": ""})
}

func UserDelete(c *gin.Context) {
	idstrs := c.QueryArray("id")
	if len(idstrs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please give the delete ids"})
		return
	}
	var ids []int
	for _, v := range idstrs {
		id, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	ret := model.UserDelete(ids)
	if ret != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ret.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": ""})
}
