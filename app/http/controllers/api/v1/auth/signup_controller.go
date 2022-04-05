package auth

import (
	v1 "api/app/http/controllers/api/v1"
	"api/app/models/user"
	"api/app/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sign *SignupController) IsPhoneExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	// 解析JSON
	if err := c.ShouldBind(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}
	//fmt.Println(request)

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
