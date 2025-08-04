// 自动生成模板TOrderAddressRecordResult
package system

import (
	"time"
)

// tOrderAddressRecordResult表 结构体  TOrderAddressRecordResult
type TOrderAddressRecordResult struct {
	Id               int64     `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                     //id字段
	FromAddressPart  string    `json:"fromAddressPart" form:"fromAddressPart" gorm:"column:from_address_part;comment:前三后四码;size:255;"` //前三后四码
	Address          string    `json:"address" form:"address" gorm:"column:address;comment:地址;size:255;"`                              //地址
	CreateTime       time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                           //创建时间
	PrivateAddress   string    `json:"privateAddress" form:"privateAddress" gorm:"column:private_address;comment:私钥;size:255;"`        //私钥
	MatchSuccessTime time.Time `json:"matchSuccessTime" form:"matchSuccessTime" gorm:"column:match_success_time;comment:匹配成功时间;"`      //匹配成功时间
}

// TableName tOrderAddressRecordResult表 TOrderAddressRecordResult自定义表名 t_order_address_record_result
func (TOrderAddressRecordResult) TableName() string {
	return "t_order_address_record_result"
}
