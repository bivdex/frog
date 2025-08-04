package system

import (
	"boost/data/server/global"
	"boost/data/server/model/system"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TMainAddrTxsService struct{}

// 创建表
func (tOrderService *TMainAddrTxsService) MigrateOrUpdateTMainAddrTxs(ctx context.Context, tableName string, tMainAddrTxs system.TMainAddrTxs) error {
	// 1. 检查并迁移表（如果不存在）
	if !global.GVA_DB_G.Migrator().HasTable(tableName) {
		if err := global.GVA_DB_G.Table(tableName).AutoMigrate(&system.TMainAddrTxs{}); err != nil {
			return fmt.Errorf("表 %s 迁移失败: %v", tableName, err)
		}
		fmt.Printf("表 %s 创建成功\n", tableName)
	}

	// 2. 查询是否存在 from_addr
	var existing system.TMainAddrTxs
	err := global.GVA_DB_G.Table(tableName).
		Where("from_addr = ?", tMainAddrTxs.FromAddr).
		First(&existing).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在：插入
			if createErr := global.GVA_DB_G.Table(tableName).Create(&tMainAddrTxs).Error; createErr != nil {
				return fmt.Errorf("插入失败: %v", createErr)
			}
			fmt.Println("插入成功")
		} else {
			// 查询出错
			return fmt.Errorf("查询失败: %v", err)
		}
	} else {
		// 存在：更新（你可以只更新部分字段）
		if existing.LatestInTime.After(tMainAddrTxs.LatestInTime) || existing.LatestInTime.Equal(tMainAddrTxs.LatestInTime) {
			return nil
		}
		fmt.Println("existing:new:newAmount", existing.LatestInTime, tMainAddrTxs.LatestInTime, tMainAddrTxs.UsdtAmount)
		existing.UsdtAmount = tMainAddrTxs.UsdtAmount
		existing.UpdatedAt = time.Now()
		existing.LatestInTime = tMainAddrTxs.LatestInTime
		if err := global.GVA_DB_G.Table(tableName).
			Where("id = ?", existing.ID).
			Updates(&existing).Error; err != nil {
			return fmt.Errorf("更新失败: %v", err)
		}
		fmt.Println("更新成功")
	}
	return nil
}

// 创建tMainAddrTxs表记录
func (tOrderService *TMainAddrTxsService) CreateOrUpdateTMainAddrTxs(ctx context.Context, tMainAddrTxs *system.TMainAddrTxs) (err error) {
	var existing system.TMainAddrTxs
	err = global.GVA_DB_G.
		Where("from_addr = ? AND main_addr = ?", tMainAddrTxs.FromAddr, tMainAddrTxs.MainAddr).
		First(&existing).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在，创建新记录
			err = global.GVA_DB_G.Create(tMainAddrTxs).Error
		} else {
			// 查询失败
			return err
		}
	} else {
		// 存在，执行更新（你可以指定更新字段）
		// if existing.LatestInTime.After(tMainAddrTxs.LatestInTime) || existing.LatestInTime.Equal(tMainAddrTxs.LatestInTime) {
		// 	return nil
		// }
		fmt.Println("existing:new:newAmount", existing.LatestInTime, tMainAddrTxs.LatestInTime, tMainAddrTxs.UsdtAmount)
		tMainAddrTxs.ID = existing.ID // 确保更新的是原记录
		tMainAddrTxs.UsdtAmount = existing.UsdtAmount
		tMainAddrTxs.UpdatedAt = time.Now()
		tMainAddrTxs.CreatedAt = existing.CreatedAt
		err = global.GVA_DB_G.Model(&existing).Updates(tMainAddrTxs).Error
	}
	return err
}

func (tOrderService *TMainAddrTxsService) GetLatestTxs(ctx context.Context, addr string) (tx system.TMainAddrTxs, canPass bool, err error) {
	err = global.GVA_DB_G.
		Where("main_addr = ?", addr).
		Order("latest_in_time DESC").
		Limit(1).
		First(&tx).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tx, true, err
		} else {
			// 其他错误
			return tx, false, err
		}
	}
	return tx, false, nil
}

func (tOrderService *TMainAddrTxsService) GetTxsList(ctx context.Context) (txs []system.TMainAddrTxs, err error) {
	err = global.GVA_DB_G.Find(&txs).Error
	return txs, err
}

func (tOrderService *TMainAddrTxsService) DeleteByFromAddr(ctx context.Context, fromAddr string) (err error) {
	err = global.GVA_DB_G.
		Where("from_addr = ?", fromAddr).
		Delete(&system.TMainAddrTxs{}).Error
	return err
}
