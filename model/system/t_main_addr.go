// 自动生成模板TMainAddr
package system

import (
	"time"
)

// tMainAddr 结构体  TMainAddr
type TMainAddr struct {
	ID            int       `gorm:"primaryKey;autoIncrement;comment:主键ID"`
	MainAddr      string    `gorm:"type:varchar(100);comment:主地址"`
	Tx24Hours     int       `gorm:"column:24_hours_txs;default:0;comment:24小时笔数"`
	TotalTxs      int       `gorm:"default:0;comment:总笔数"`
	AvgUsdtAmount float64   `gorm:"type:decimal(64,2);comment:平均金额"`
	CreatedAt     time.Time `gorm:"type:datetime;comment:创建时间"`
	UpdatedAt     time.Time `gorm:"type:datetime;comment:更新时间"`
}

// TableName tMainAddr TMainAddr t_main_addr
func (TMainAddr) TableName() string {
	return "t_main_addr"
}
