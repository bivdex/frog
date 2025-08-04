package system

import (
	"boost/data/server/global"
	"boost/data/server/model/system"
	"context"
	"errors"
	"gorm.io/gorm"
)

type TAddressActivitydService struct{}

// CreateTAddressActivity 创建TAddressActivity表记录
func (tAddressActivitydService *TAddressActivitydService) CreateTAddressActivity(ctx context.Context, tAddressActivity *system.TAddressActivity) (err error) {
	err = global.GVA_DB_Local.Create(tAddressActivity).Error
	return err
}

// DeleteTAddressActivity 删除TAddressActivity表记录
func (tAddressActivitydService *TAddressActivitydService) DeleteTAddressActivity(ctx context.Context, id string) (err error) {
	err = global.GVA_DB_Local.Delete(&system.TAddressActivity{}, "id = ?", id).Error
	return err
}

// UpdateTAddressActivity 更新TAddressActivity表记录
func (tAddressActivitydService *TAddressActivitydService) UpdateTAddressActivity(ctx context.Context, tAddressActivity system.TAddressActivity) (err error) {
	err = global.GVA_DB_Local.Model(&system.TAddressActivity{}).Where("id = ?", tAddressActivity.Id).Updates(&tAddressActivity).Error
	return err
}

// GetTAddressActivity 根据id获取TAddressActivity表记录
func (tAddressActivitydService *TAddressActivitydService) GetTAddressActivityByID(ctx context.Context, id string) (tAddressActivity system.TAddressActivity, err error) {
	err = global.GVA_DB_Local.Where("id = ?", id).First(&tAddressActivity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tAddressActivity, nil
	} else if err != nil {
		return tAddressActivity, err
	}
	return
}

// GetTAddressActivity
func (tAddressActivitydService *TAddressActivitydService) GetTAddressActivity(ctx context.Context, _activeAddress, _regularAddress string) (tAddressActivity system.TAddressActivity, err error) {
	err = global.GVA_DB_Local.Where(" active_address = ? AND  regular_address = ?  ", _activeAddress, _regularAddress).First(&tAddressActivity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tAddressActivity, nil
	} else if err != nil {
		return tAddressActivity, err
	}
	return
}

// GetTOrderInfoList 分页获取tOrder表记录
func (tAddressActivitydService *TAddressActivitydService) GetFullTOrderInfoList(ctx context.Context) (list []system.TAddressActivity, total int64, err error) {

	db := global.GVA_DB_Local.Model(&system.TAddressActivity{})
	var tOrders []system.TAddressActivity
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&tOrders).Error
	return tOrders, total, err
}
