// 自动生成模板TReceiveOrder
package system

import (
	"time"
)

// t_order_from_address 结构体  TOrderFromAddress
type TOrderFromAddress struct {
	Id          int64     `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`
	FromAddress string    `json:"fromAddress" form:"fromAddress" gorm:"column:from_address;comment:from地址;size:255;"` //from地址
	LastAmount  float64   `json:"lastAmount" form:"lastAmount" gorm:"column:last_amount;comment:最后到账金额;size:64;"`     //最后到账金额
	Balance     float64   `json:"balance" form:"balance" gorm:"column:balance;comment:余额;size:64;"`                   //余额
	LastTime    time.Time `json:"lastTime" form:"lastTime" gorm:"column:last_time;comment:最后到账时间;"`                   //创建时间
	QueryTime   time.Time `json:"queryTime" form:"queryTime" gorm:"column:query_time;comment:查询时间;"`                  //创建时间
	CreateTime  time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`               //创建时间
}

// TableName tOrderFromAddress表 TOrderFromAddress自定义表名 t_order_from_address
func (TOrderFromAddress) TableName() string {
	return "t_order_from_address"
}
