// 自动生成模板TOrderTransferFromAddress
package system

import (
	"time"
)

// tOrderTransferFromAddress表 结构体  TOrderTransferFromAddress
type TOrderTransferFromAddress struct {
	Id              int64     `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`                                           //id字段
	OrderNo         string    `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;size:100;"`                                  //订单号
	FromAddressPart string    `json:"fromAddressPart" form:"fromAddressPart" gorm:"column:from_address_part;comment:充值地址前3位和后4位;size:255;"` //充值地址前3位和后4位
	FromAddress     string    `json:"fromAddress" form:"fromAddress" gorm:"column:from_address;comment:充值地址;size:255;"`                     //充值地址
	ToAddress       string    `json:"toAddress" form:"toAddress" gorm:"column:to_address;comment:接收地址;size:255;"`                           //接收地址
	Amount          float64   `json:"amount" form:"amount" gorm:"column:amount;comment:数量;size:64;"`                                        //数量
	PrivateKey      string    `json:"privateKey" form:"privateKey" gorm:"column:private_key;comment:私钥;size:255;"`                          //私钥
	Status          int       `json:"status" form:"status" gorm:"column:status;comment:状态：1.转账中 2.以转账;size:10;"`                            //状态：1.转账中 2.以转账
	OrderTime       string    `json:"orderTime" form:"orderTime" gorm:"column:order_time;comment:订单时间;size:255;"`                           //订单时间
	CreateTime      time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                                 //创建时间
	TransferHashTx  string    `json:"transferHashTx" form:"transferHashTx" gorm:"column:transfer_hash_tx;comment:交易hash;size:255;"`         //交易hash
	TransferTime    time.Time `json:"transferTime" form:"transferTime" gorm:"column:transfer_time;comment:转账时间;"`                           //转账时间
}

// TableName tOrderTransferFromAddress表 TOrderTransferFromAddress自定义表名 t_order_transfer_from_address
func (TOrderTransferFromAddress) TableName() string {
	return "t_order_transfer_from_address"
}
