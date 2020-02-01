package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"hfga.com.cn/ginorm/controller"
	"strings"
)

func AuthRequired(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth != "1234567" {
		c.JSON(401, gin.H{
			"message": "not authorized",
		})
		c.Abort()
	}
	c.Set("userName", "haoyun") //别的地方可以用
}

func AuthReq(c *gin.Context) {

	token, err := request.ParseFromRequest(c.Request, authorizationHeaderExtractor, //request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			// welcome to gin 这段也在加密的时候用到了
			return []byte("welcome to gin"), nil
		})

	if err == nil {
		if token.Valid {
			//TODO 这地方实现了解析username,然后可以继续做别的事情
			claim := token.Claims.(jwt.MapClaims)
			// 设置userName，后续可能用。
			c.Set("userName", claim["username"].(string))
			//val, _ := c.Get("userName")
			//ctx.Input.Data()["username"] = claim["username"].(string)
			// 下面是根据用户名获取用户名有哪些权限，或者当请求来时，看看该用户有没有此权限
			//erra, authority := getUserAuthority(ctx, claim["username"].(string))
			//if erra != nil {
			//	if strings.Contains(erra.Error(), "not found"){
			//		authRespose(ctx, false, "Unauthorized", 401)
			//		return
			//	}
			//	authRespose(ctx, false, fmt.Sprintf("server occured some err, err: %s", erra.Error()), 500)
			//	return
			//}
			//if authority != "system-admin" {
			//	authRespose(ctx, false, "you don't have the right to do so", 401)
			//}
		} else {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		}
	} else {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
	}
}

// Strips 'Bearer ' or 'jwt' prefix from bearer token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 3 && strings.ToUpper(tok[0:4]) == "JWT " {
		return tok[4:], nil
	}
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

//func authRespose(ctx *context.Context, status bool, errMsg string, statusCode int) {
//	replyStatus := entity.ReplyStatus{status, errMsg}
//	ctx.Output.Status = statusCode
//	ctx.Output.JSON(replyStatus, false, true)
//}

// Extract bearer token from Authorization header
// Uses PostExtractionFilter to strip "Bearer " prefix from header
var authorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter: stripBearerPrefixFromTokenString,
}

func Route(r *gin.Engine) {
	//auth
	r.POST("/gin/auth", controller.AuthPost)
	v1 := r.Group("/gin")
	//v1.Use(AuthRequired)
	v1.Use(AuthReq)
	v1.POST("/roles", controller.RolePost)
	v1.GET("/roles", controller.RoleGet)
	v1.PUT("/roles", controller.RolePut)
	v1.DELETE("/roles", controller.RoleDelete)
	// perms
	v1.POST("/perms", controller.PermPost)
	v1.GET("/perms", controller.PermGet)
	v1.PUT("/perms", controller.PermPut)
	v1.DELETE("/perms", controller.PermDelete)
	// 用户
	v1.POST("/users", controller.UserPost)
	v1.GET("/users", controller.UserGet)
	v1.PUT("/users", controller.UserPut)
	v1.DELETE("/users", controller.UserDelete)
}
