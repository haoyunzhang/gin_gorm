package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hfga.com.cn/ginorm/model"
	"net/http"
	"time"
)

func AuthPost(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ret, err := model.UserGetOne(user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	genPdUser := model.User{
		Name: user.Name,
		PassWord: user.PassWord,
	}
	if ret.PassWord != model.GenerateId(genPdUser) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is not correct"})
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// 过期时间
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["username"] = user.Name
	token.Claims = claims

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// welcome to gin 这段也在解码的时候用到了
	tokenString, err := token.SignedString([]byte("welcome to gin"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "error": ""})
}
