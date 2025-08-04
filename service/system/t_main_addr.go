package system

import (
	"boost/data/server/global"
	"boost/data/server/model/system"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type TMainAddrService struct{}

// CreateTMainAddr 创建tMainAddr表记录
func (tOrderService *TMainAddrService) CreateOrUpdateTMainAddr(ctx context.Context, tMainAddr *system.TMainAddr) error {

	var existing system.TMainAddr
	err := global.GVA_DB_G.Where("main_addr = ?", tMainAddr.MainAddr).First(&existing).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在：创建
			return global.GVA_DB_G.Create(tMainAddr).Error
		}
		// 查询出错
		return err
	}

	// 存在：更新
	tMainAddr.ID = existing.ID // 确保更新目标是这条记录
	tMainAddr.UpdatedAt = time.Now()
	tMainAddr.CreatedAt = existing.CreatedAt
	return global.GVA_DB_G.Model(&existing).Updates(tMainAddr).Error
}

func (tOrderService *TMainAddrService) GetTMainAddr(ctx context.Context, tMainAddr *system.TMainAddr) error {

	var existing system.TMainAddr
	err := global.GVA_DB_G.Where("main_addr = ?", tMainAddr.MainAddr).First(&existing).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在：创建
			return global.GVA_DB_G.Create(tMainAddr).Error
		}
		// 查询出错
		return err
	}

	// 存在：更新
	tMainAddr.ID = existing.ID // 确保更新目标是这条记录
	tMainAddr.UpdatedAt = time.Now()
	tMainAddr.CreatedAt = existing.CreatedAt
	return global.GVA_DB_G.Model(&existing).Updates(tMainAddr).Error
}

func (tOrderService *TMainAddrService) GetTxsList(ctx context.Context) (txs []system.TMainAddr, err error) {
	err = global.GVA_DB_G.Find(&txs).Error
	return txs, err
}
