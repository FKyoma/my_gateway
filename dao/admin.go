package dao

import (
	"github.com/e421083458/go_gateway_demo/dto"
	"github.com/e421083458/go_gateway_demo/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string    `json:"area_name" gorm:"column:area_name" description:"管理员用户名"`
	Salt      string       `json:"city_id" gorm:"column:city_id" description:"盐"`
	Password  string     `json:"user_id" gorm:"column:user_id" description:"密码"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  int 	`json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (t *Admin) TableName() string {
	return "gateway_admin"
}

func (t *Admin)LoginCheck(c *gin.Context,tx *gorm.DB,param *dto.AdminLoginInput)(*Admin,error)  {
	adminInfo,err:=t.Find(c,tx,(&Admin{UserName: param.Username,IsDelete: 0}))
	if err!=nil{
		return nil,errors.New("用户信息不存在")
	}
	param.Password
	adminInfo.Salt

	return nil,nil
}

func (t *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out:=&Admin{}
	err:=tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	//err := tx.SetCtx(c).Where(search).Find(area).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
