// 自动生成模板TReceiveOrder
package system

import (
	"time"
)

// TAddressActivity 结构体  TAddressActivity
type TAddressActivity struct {
	//id字段
	Id                    int64     `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`
	ActiveAddress         string    `json:"active_address" form:"active_address" gorm:"column:active_address;comment:活跃地址;size:100;"`     //订单号
	RegularAddress        string    `json:"regular_address" form:"regular_address" gorm:"column:regular_address;comment:常转出地址;size:255;"` //目标地址
	AccountBalance        string    `json:"account_balance" form:"account_balance" gorm:"column:account_balance;comment:当前余额;size:64;"`
	CurrentOutboundAmount string    `json:"current_outbound_amount" form:"current_outbound_amount" gorm:"column:current_outbound_amount;comment:当前余额;size:64;"`
	TotalOutboundAmount   string    `json:"total_outbound_amount" form:"total_outbound_amount" gorm:"column:total_outbound_amount;comment:转出总额;size:64;"`
	AverageAmount         string    `json:"average_amount" form:"average_amount" gorm:"column:average_amount;comment:转出总额;size:64;"`
	RecentThreeAvg        string    `json:"recent_three_avg" form:"recent_three_avg" gorm:"column:recent_three_avg;comment:转出总额;size:64;"`
	OutboundCount         int       `json:"outbound_count" form:"outbound_count" gorm:"column:outbound_count;comment:outbound_count;"`                                 //cztimes
	CreatedTime           time.Time `json:"created_time" form:"created_time" gorm:"column:created_time;comment:创建时间;"`                                                 //创建时间
	LastTransactionTime   time.Time `json:"last_transaction_time" form:"last_transaction_time" gorm:"column:last_transaction_time;comment:创建时间;"`                      //创建时间
	CreatedBlock          int64     `json:"created_block" form:"created_block" gorm:"column:created_block;comment:created_block;"`                                     //cztimes
	LastTransactionBlock  int64     `json:"last_transaction_block" form:"last_transaction_block" gorm:"column:last_transaction_block;comment:last_transaction_block;"` //cztimes

}

// TableName TAddressActivity  t_address_activity
func (TAddressActivity) TableName() string {
	return "t_address_activity"
}
