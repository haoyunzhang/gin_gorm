package controller

import (
	"github.com/gin-gonic/gin"
	"hfga.com.cn/ginorm/model"
	"net/http"
	"strconv"
)

func RolePost(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := model.RolePost(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"error": ""})
}

func RoleGet(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		roleGetAll(c)
		return
	}
	roleGetOne(c, name)
}

func roleGetOne(c *gin.Context, name string) {
	ret, err := model.RoleGetOne(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ret, "error": ""})
}

func roleGetAll(c *gin.Context) {
	ret, err := model.RoleGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ret, "error": ""})
}

func RolePut(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := model.RolePut(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"error": ""})
}

func RoleDelete(c *gin.Context) {
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
	ret := model.RoleDelete(ids)
	if ret != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ret.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": ""})
}
