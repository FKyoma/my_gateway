package dao

import (
	"github.com/e421083458/go_gateway_demo/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type ServiceInfo struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	ServiceName string    `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	LoadType    int       `json:"load_type" gorm:"column:load_type" description:"负载类型 0=http 1=tcp 2=grpc"`
	UpdatedAt   time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt   time.Time `json:"create_at" gorm:"column:create_at" description:"添加时间"`
	IsDelete    int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除; 0:否 1:是"`
}

func (t *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

func (t *ServiceInfo)PageList(c *gin.Context,tx *gorm.DB,param *dto.ServiceListInput)([]ServiceInfo,int64,error){
	total:=0
	list:=[]ServiceInfo{}
	query:=tx.WithContext(c)

	if param.Info!=""{
		query=query.Where("(service_name like %?% or service_desc like %?%)",param.Info,param.Info)
	}

}

func (t *ServiceInfo) Find(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out := &ServiceInfo{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	//err := tx.SetCtx(c).Where(search).Find(area).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (t *ServiceInfo) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(t).Error
}
