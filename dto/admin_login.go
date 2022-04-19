package dto

import (
	"github.com/e421083458/go_gateway_demo/public"
	"github.com/gin-gonic/gin"
	"time"
)

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type AdminLoginInput struct {
	Username string `form:"username" json:"username" comment:"用户名"  validate:"required,valid_username" example:"admin"` //管理员用户名
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:"123456"`                  //密码
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct {
	Token string `form:"token" json:"token" comment:"token"  validate:"" example:"token"` //token

}
