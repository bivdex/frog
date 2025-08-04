package system

import (
	"boost/data/server/global"
	"boost/data/server/model/system"
	systemReq "boost/data/server/model/system/request"
	"context"
	"errors"
	"gorm.io/gorm"
)

type TOrderFromAddressService struct{}

// CreateTOrder 创建tOrder表记录
func (tOrderFromAddressService *TOrderFromAddressService) CreateTOrderFromAddress(ctx context.Context, tOrder *system.TOrderFromAddress) (err error) {
	err = global.GVA_DB_E.Create(tOrder).Error
	return err
}

// DeleteTOrder 删除tOrder表记录
func (tOrderFromAddressService *TOrderFromAddressService) DeleteTOrderFromAddress(ctx context.Context, id string) (err error) {
	err = global.GVA_DB_E.Delete(&system.TOrderFromAddress{}, "id = ?", id).Error
	return err
}

// DeleteTOrderByIds 批量删除tOrder表记录
func (tOrderFromAddressService *TOrderFromAddressService) DeleteTOrderByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB_E.Delete(&[]system.TOrderFromAddress{}, "id in ?", ids).Error
	return err
}

// UpdateTOrder 更新tOrder表记录
func (tOrderFromAddressService *TOrderFromAddressService) UpdateTOrderFromAddress(ctx context.Context, tOrder system.TOrderFromAddress) (err error) {
	err = global.GVA_DB_E.Model(&system.TOrderFromAddress{}).Where("id = ?", tOrder.Id).Updates(&tOrder).Error
	return err
}

// GetTOrder 根据id获取tOrder表记录
func (tOrderFromAddressService *TOrderFromAddressService) GetTOrderFromAddress(ctx context.Context, id string) (tOrder system.TOrderFromAddress, err error) {
	err = global.GVA_DB_E.Where("id = ?", id).First(&tOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tOrder, nil
	} else if err != nil {
		return tOrder, err
	}
	return
}

// GetTOrder 根据id获取tOrder表记录
func (tOrderFromAddressService *TOrderFromAddressService) GetTOrderFromAddressByFromAddress(ctx context.Context, from_address string) (tOrder system.TOrderFromAddress, err error) {
	err = global.GVA_DB_E.Where("from_address = ?", from_address).First(&tOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tOrder, nil
	} else if err != nil {
		return tOrder, err
	}

	return
}

// GetTOrderInfoList 分页获取tOrder表记录
func (tOrderFromAddressService *TOrderFromAddressService) GetTOrderInfoList(ctx context.Context, info systemReq.TOrderSearch) (list []system.TOrderFromAddress, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB_E.Model(&system.TOrderFromAddress{})
	var tOrders []system.TOrderFromAddress
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tOrders).Error
	return tOrders, total, err
}

// GetTOrderInfoList 分页获取tOrder表记录
func (tOrderFromAddressService *TOrderFromAddressService) GetFullTOrderInfoList(ctx context.Context) (list []system.TOrderFromAddress, total int64, err error) {

	db := global.GVA_DB_E.Model(&system.TOrderFromAddress{})
	var tOrders []system.TOrderFromAddress
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&tOrders).Error
	return tOrders, total, err
}
