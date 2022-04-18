package dao

type AccessControl struct {
	ID                int64  `json:"id" gorm:"primary_key"`
	ServiceID         int64  `json:"id" gorm:"primary_key"`
	OpenAuth          int    `json:"oepn_auth" gorm:"column:oepn_auth" description:"是否开启权限"`
	BlackList         string `json:"clientip_flow_limit" gorm:"column:clientip_flow_limit" description:"客户端ip限流"`
	WhiteList         string `json:"clientip_flow_limit" gorm:"column:clientip_flow_limit" description:"客户端ip限流"`
	WhiteHostName     string `json:"clientip_flow_limit" gorm:"column:clientip_flow_limit" description:"客户端ip限流"`
	ClientIPFlowLimit int64  `json:"clientip_flow_limit" gorm:"column:clientip_flow_limit" description:"客户端ip限流"`
	ServiceFlowLimit  int64  `json:"service_flow_limit" gorm:"column:service_flow_limit" description:"服务端限流"`
}
