package controller

import (
	"github.com/gin-gonic/gin"
	"hfga.com.cn/ginorm/model"
	"net/http"
	"strconv"
)

func PermPost(c *gin.Context) {
	var perm model.Perm
	if err := c.ShouldBindJSON(&perm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := model.PermPost(&perm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"error": ""})
}

func PermGet(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		permGetAll(c)
		return
	}
	permGetOne(c, name)
}

func permGetOne(c *gin.Context, name string) {
	ret, err := model.PermGetOne(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ret, "error": ""})
}

func permGetAll(c *gin.Context) {
	ret, err := model.PermGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ret, "error": ""})
}

func PermPut(c *gin.Context) {
	var perm model.Perm
	if err := c.ShouldBindJSON(&perm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := model.PermPut(&perm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"error": ""})
}

func PermDelete(c *gin.Context) {
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
	ret := model.PermDelete(ids)
	if ret != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ret.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": ""})
}
