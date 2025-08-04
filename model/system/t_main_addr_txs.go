// 自动生成模板TMainAddr
package system

import (
	"time"
)

// tMainAddrTxs 结构体  TMainAddrTxs
type TMainAddrTxs struct {
	ID           int       `gorm:"primaryKey;autoIncrement;comment:主键ID"`
	MainAddr     string    `gorm:"type:varchar(100);comment:主地址"`
	FromAddr     string    `gorm:"type:varchar(100);comment:from地址"`
	UsdtAmount   float64   `gorm:"type:decimal(64,2);comment:金额"`
	CreatedAt    time.Time `gorm:"type:datetime;comment:创建时间"`
	UpdatedAt    time.Time `gorm:"type:datetime;comment:更新时间"`
	LatestInTime time.Time `gorm:"type:datetime;comment:最后入账时间"`
}

// TableName tMainAddr TMainAddr t_main_addr
func (TMainAddrTxs) TableName() string {
	return "t_main_addr_txs"
}
