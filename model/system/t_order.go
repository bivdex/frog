// 自动生成模板TOrder
package system

import (
	"time"
)

// tOrder表 结构体  TOrder
type TOrder struct {
	Id              int64     `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                           //id字段
	OrderNo         string    `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;size:100;"`                                  //订单号
	FromAddressPart string    `json:"fromAddressPart" form:"fromAddressPart" gorm:"column:from_address_part;comment:充值地址前3位和后4位;size:255;"` //充值地址前3位和后4位
	FromAddress     string    `json:"fromAddress" form:"fromAddress" gorm:"column:from_address;comment:充值地址;size:255;"`                     //充值地址
	ToAddress       string    `json:"toAddress" form:"toAddress" gorm:"column:to_address;comment:接收地址;size:255;"`                           //接收地址
	Amount          float64   `json:"amount" form:"amount" gorm:"column:amount;comment:数量;size:64;"`                                        //数量
	PrivateKey      string    `json:"privateKey" form:"privateKey" gorm:"column:private_key;comment:私钥;size:255;"`                          //私钥
	TransferType    bool      `json:"transferType" form:"transferType" gorm:"column:transfer_type;comment:转账类型;"`                           //转账类型
	Status          int       `json:"status" form:"status" gorm:"column:status;comment:状态，0 初始化，1 处理中 2等待转出方地址激活 9已完成;size:10;"`            //状态，0 初始化，1 处理中 2等待转出方地址激活 9已完成
	OrderTime       string    `json:"orderTime" form:"orderTime" gorm:"column:order_time;comment:订单时间;size:255;"`                           //订单时间
	MatchTime       time.Time `json:"matchTime" form:"matchTime" gorm:"column:match_time;comment:匹配时间;"`                                    //匹配时间
	CreateTime      time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                                 //创建时间
	TransferTime    time.Time `json:"transferTime" form:"transferTime" gorm:"column:transfer_time;comment:转账时间;"`                           //转账时间
}

// TableName tOrder表 TOrder自定义表名 t_order
func (TOrder) TableName() string {
	return "t_order"
}
