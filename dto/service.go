package dto

import (
	"github.com/e421083458/go_gateway_demo/public"
	"github.com/gin-gonic/gin"
)

type ServiceListInput struct {
	Info     string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`                      //关键词
	PageNo   int    `form:"page_no" json:"page_no" comment:"页数"   validate:"required" example:"1"`        //页数
	PageSize int    `form:"page_size" json:"page_size" comment:"每页页数"   validate:"required" example:"20"` //每页页数
}

func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceListItemOutput struct {
	ID          int64  `json:"id" form:"id"`                     //id
	ServiceName string `json:"service_name" form:"service_name"` //服务名称
	ServiceDesc string `json:"service_desc" form:"service_desc"` //服务描述
	LoadType    int    `json:"load_type" form:"load_type"`       //类型
	ServiceAddr string `json:"service_addr" form:"service_addr"` //服务地址
	Qps         int64  `json:"qps" form:"qps"`                   //qps
	Qpd         int64  `json:"qpd" form:"qpd"`                   //qpd
	TotalNode   int    `json:"total_node" form:"total_node"`     //节点数
}

type ServiceListOutput struct {
	Total int64  `form:"total" json:"total" comment:"总数"   validate:"" example:""`        //总数
	List  string `form:"list" json:"list" comment:"列表"   validate:"required" example:"1"` //列表
}
