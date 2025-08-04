package system

import (
	"boost/data/server/global"
	"boost/data/server/model/system"
	systemReq "boost/data/server/model/system/request"
	"context"
	"errors"
	"gorm.io/gorm"
)

type TOrderAddressRecordResultService struct{}

// CreateTOrderAddressRecordResult 创建tOrderAddressRecordResult表记录
func (tOrderAddressRecordResultService *TOrderAddressRecordResultService) CreateTOrderAddressRecordResult(ctx context.Context, tOrderAddressRecordResult *system.TOrderAddressRecordResult) (err error) {
	err = global.GVA_DB.Create(tOrderAddressRecordResult).Error
	return err
}

// DeleteTOrderAddressRecordResult 删除tOrderAddressRecordResult表记录
func (tOrderAddressRecordResultService *TOrderAddressRecordResultService) DeleteTOrderAddressRecordResult(ctx context.Context, id int64) (err error) {
	err = global.GVA_DB.Delete(&system.TOrderAddressRecordResult{}, "id = ?", id).Error
	return err
}

// DeleteTOrderAddressRecordResultByIds 批量删除tOrderAddressRecordResult表记录
func (tOrderAddressRecordResultService *TOrderAddressRecordResultService) DeleteTOrderAddressRecordResultByIds(ctx context.Context, ids []int64) (err error) {
	err = global.GVA_DB.Delete(&[]system.TOrderAddressRecordResult{}, "id in ?", ids).Error
	return err
}

// UpdateTOrderAddressRecordResult 更新tOrderAddressRecordResult表记录
func (tOrderAddressRecordResultService *TOrderAddressRecordResultService) UpdateTOrderAddressRecordResult(ctx context.Context, tOrderAddressRecordResult system.TOrderAddressRecordResult) (err error) {
	err = global.GVA_DB.Model(&system.TOrderAddressRecordResult{}).Where("id = ?", tOrderAddressRecordResult.Id).Updates(&tOrderAddressRecordResult).Error
	return err
}

// GetTOrderAddressRecordResult 根据id获取tOrderAddressRecordResult表记录
func (tOrderAddressRecordResultService *TOrderAddressRecordResultService) GetTOrderAddressRecordResult(ctx context.Context, id int64) (tOrderAddressRecordResult system.TOrderAddressRecordResult, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tOrderAddressRecordResult).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tOrderAddressRecordResult, nil
	} else if err != nil {
		return tOrderAddressRecordResult, err
	}

	return
}

// GetTOrderAddressRecordResult 根据id获取tOrderAddressRecordResult表记录
func (tOrderAddressRecordResultService *TOrderAddressRecordResultService) GetTOrderAddressRecordResultByFromAddressPart(ctx context.Context, fromAddressPart string) (tOrderAddressRecordResult system.TOrderAddressRecordResult, err error) {
	err = global.GVA_DB.Where("from_address_part = ?", fromAddressPart).First(&tOrderAddressRecordResult).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tOrderAddressRecordResult, nil
	} else if err != nil {
		return tOrderAddressRecordResult, err
	}

	return
}

// GetTOrderAddressRecordResultInfoList 分页获取tOrderAddressRecordResult表记录
func (tOrderAddressRecordResultService *TOrderAddressRecordResultService) GetTOrderAddressRecordResultInfoList(ctx context.Context, info systemReq.TOrderAddressRecordResultSearch) (list []system.TOrderAddressRecordResult, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.TOrderAddressRecordResult{})
	var tOrderAddressRecordResults []system.TOrderAddressRecordResult
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tOrderAddressRecordResults).Error
	return tOrderAddressRecordResults, total, err
}
func (tOrderAddressRecordResultService *TOrderAddressRecordResultService) GetTOrderAddressRecordResultPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
