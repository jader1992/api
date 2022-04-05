package routes

import (
	"api/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAPIRouters RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRouters(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "world!",
			})
		})
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机号码是否已被注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}
}
