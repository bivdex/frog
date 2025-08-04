// 自动生成模板TOrderToAddressRecord
package system

import (
	"time"
)

// tOrderToAddressRecord表 结构体  TOrderToAddressRecord
type TOrderToAddressRecord struct {
	Id              int64     `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                     //id字段
	FromAddressPart string    `json:"fromAddressPart" form:"fromAddressPart" gorm:"column:from_address_part;comment:前三后四码;size:255;"` //前三后四码
	ToAddress       string    `json:"toAddress" form:"toAddress" gorm:"column:to_address;comment:目标地址;size:255;"`                     //目标地址
	CreateTime      time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                           //创建时间
}

// TableName tOrderToAddressRecord表 TOrderToAddressRecord自定义表名 t_order_to_address_record
func (TOrderToAddressRecord) TableName() string {
	return "t_order_to_address_record"
}
