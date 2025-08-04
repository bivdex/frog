package system

import (
	"boost/data/server/global"
	"boost/data/server/model/system"
	systemReq "boost/data/server/model/system/request"
	"context"
	"errors"
	"gorm.io/gorm"
)

type TOrderToAddressRecordService struct{}

// CreateTOrderToAddressRecord 创建tOrderToAddressRecord表记录
func (tOrderToAddressRecordService *TOrderToAddressRecordService) CreateTOrderToAddressRecord(ctx context.Context, tOrderToAddressRecord *system.TOrderToAddressRecord) (err error) {
	err = global.GVA_DB_D.Create(tOrderToAddressRecord).Error
	return err
}

// DeleteTOrderToAddressRecord 删除tOrderToAddressRecord表记录
func (tOrderToAddressRecordService *TOrderToAddressRecordService) DeleteTOrderToAddressRecord(ctx context.Context, id string) (err error) {
	err = global.GVA_DB_D.Delete(&system.TOrderToAddressRecord{}, "id = ?", id).Error
	return err
}

// DeleteTOrderToAddressRecordByIds 批量删除tOrderToAddressRecord表记录
func (tOrderToAddressRecordService *TOrderToAddressRecordService) DeleteTOrderToAddressRecordByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB_D.Delete(&[]system.TOrderToAddressRecord{}, "id in ?", ids).Error
	return err
}

// UpdateTOrderToAddressRecord 更新tOrderToAddressRecord表记录
func (tOrderToAddressRecordService *TOrderToAddressRecordService) UpdateTOrderToAddressRecord(ctx context.Context, tOrderToAddressRecord system.TOrderToAddressRecord) (err error) {
	err = global.GVA_DB_D.Model(&system.TOrderToAddressRecord{}).Where("id = ?", tOrderToAddressRecord.Id).Updates(&tOrderToAddressRecord).Error
	return err
}

// GetTOrderToAddressRecord 根据id获取tOrderToAddressRecord表记录
func (tOrderToAddressRecordService *TOrderToAddressRecordService) GetTOrderToAddressRecordByID(ctx context.Context, id string) (tOrderToAddressRecord system.TOrderToAddressRecord, err error) {
	err = global.GVA_DB_D.Where("id = ?", id).First(&tOrderToAddressRecord).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tOrderToAddressRecord, nil
	} else if err != nil {
		return tOrderToAddressRecord, err
	}

	return
}

// GetTOrderToAddressRecord 根据id获取tOrderToAddressRecord表记录 根据活跃地址和前三后四码都相同则写入C库
func (tOrderToAddressRecordService *TOrderToAddressRecordService) GetTOrderToAddressRecord(ctx context.Context, toAddress, fromAddressPart string) (tOrderToAddressRecord system.TOrderToAddressRecord, err error) {
	err = global.GVA_DB_D.Where(" to_address = ? AND  from_address_part = ?  ", toAddress, fromAddressPart).First(&tOrderToAddressRecord).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tOrderToAddressRecord, nil
	} else if err != nil {
		return tOrderToAddressRecord, err
	}

	return
}

// GetTOrderToAddressRecordInfoList 分页获取tOrderToAddressRecord表记录
func (tOrderToAddressRecordService *TOrderToAddressRecordService) GetTOrderToAddressRecordInfoList(ctx context.Context, info systemReq.TOrderToAddressRecordSearch) (list []system.TOrderToAddressRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB_D.Model(&system.TOrderToAddressRecord{})
	var tOrderToAddressRecords []system.TOrderToAddressRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tOrderToAddressRecords).Error
	return tOrderToAddressRecords, total, err
}
func (tOrderToAddressRecordService *TOrderToAddressRecordService) GetTOrderToAddressRecordPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
