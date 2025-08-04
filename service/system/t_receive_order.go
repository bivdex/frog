package system

import (
	"boost/data/server/global"
	"boost/data/server/model/system"
	systemReq "boost/data/server/model/system/request"
	"context"
	"errors"
	"gorm.io/gorm"
)

type TReceiveOrderService struct{}

// CreateTReceiveOrder 创建tReceiveOrder表记录
func (tReceiveOrderService *TReceiveOrderService) CreateTReceiveOrder(ctx context.Context, tReceiveOrder *system.TReceiveOrder) (err error) {
	err = global.GVA_DB_C.Create(tReceiveOrder).Error
	return err
}

// DeleteTReceiveOrder 删除tReceiveOrder表记录
//func (tReceiveOrderService *TReceiveOrderService) DeleteTReceiveOrder(ctx context.Context, tReceiveOrder system.TReceiveOrder) (err error) {
//	err = global.GVA_DB_C.Model(&system.TReceiveOrder{}).Where("id = ?", tReceiveOrder.ID).Updates(&tReceiveOrder).Error
//	return err
//}
//
////// DeleteTReceiveOrderByIds 批量删除tReceiveOrder表记录
////func (tReceiveOrderService *TReceiveOrderService) DeleteTReceiveOrderByIds(ctx context.Context, ids []string) (err error) {
////	err = global.GVA_DB.Delete(&[]system.TReceiveOrder{}, "id in ?", ids).Error
////	return err
////}
//
//// UpdateTReceiveOrder 更新tReceiveOrder表记录
//func (tReceiveOrderService *TReceiveOrderService) UpdateTReceiveOrder(ctx context.Context, tReceiveOrder system.TReceiveOrder) (err error) {
//	err = global.GVA_DB_C.Model(&system.TReceiveOrder{}).Where("id = ?", tReceiveOrder.ID).Updates(&tReceiveOrder).Error
//	return err
//}

// GetTReceiveOrder 根据id获取tReceiveOrder表记录
func (tReceiveOrderService *TReceiveOrderService) GetTReceiveOrderByOrderNo(ctx context.Context, order_no string) (tReceiveOrder system.TReceiveOrder, err error) {
	err = global.GVA_DB_C.Where("order_no = ?", order_no).Where("is_del = ?", "0").First(&tReceiveOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tReceiveOrder, nil
	} else if err != nil {
		return tReceiveOrder, err
	}
	return
}

// GetTReceiveOrder 根据id获取tReceiveOrder表记录
func (tReceiveOrderService *TReceiveOrderService) GetTReceiveOrder(ctx context.Context, id string) (tReceiveOrder system.TReceiveOrder, err error) {
	err = global.GVA_DB_C.Where("id = ?", id).Where("is_del = ?", "0").First(&tReceiveOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tReceiveOrder, nil
	} else if err != nil {
		return tReceiveOrder, err
	}
	return
}

// UpdateTReceiveOrder 更新tReceiveOrder表记录
func (tReceiveOrderService *TReceiveOrderService) UpdateTReceiveOrderByFromAddressPart(ctx context.Context, tReceiveOrder system.TReceiveOrder) (err error) {
	err = global.GVA_DB_C.Model(&system.TReceiveOrder{}).Where("from_address_part = ?", tReceiveOrder.FromAddressPart).Updates(&tReceiveOrder).Error
	return err
}

// GetTReceiveOrder 根据id获取tReceiveOrder表记录
func (tReceiveOrderService *TReceiveOrderService) GetTReceiveOrderByFromAddressPartAndToAddress(ctx context.Context, _fromAddressPart, _toAddress string) (list []system.TReceiveOrder, total int64, err error) {
	//err = global.GVA_DB_C.Where("from_address_part = ?", _fromAddressPart).Where("is_del = ?", "0").First(&tReceiveOrder).Error

	// 创建db
	db := global.GVA_DB_C.Model(&system.TReceiveOrder{}).Where("from_address_part = ? and to_address = ?", _fromAddressPart, _toAddress).Where("is_del = ?", "0")
	var tReceiveOrders []system.TReceiveOrder

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&tReceiveOrders).Error
	return tReceiveOrders, total, err

	return
}

// GetTReceiveOrder 根据id获取tReceiveOrder表记录
func (tReceiveOrderService *TReceiveOrderService) GetTReceiveOrderByFromAddressPart(ctx context.Context, _fromAddressPart string) (list []system.TReceiveOrder, total int64, err error) {
	//err = global.GVA_DB_C.Where("from_address_part = ?", _fromAddressPart).Where("is_del = ?", "0").First(&tReceiveOrder).Error

	// 创建db
	db := global.GVA_DB_C.Model(&system.TReceiveOrder{}).Where("from_address_part = ?", _fromAddressPart).Where("is_del = ?", "0")
	var tReceiveOrders []system.TReceiveOrder

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&tReceiveOrders).Error
	return tReceiveOrders, total, err

	return
}

// GetTReceiveOrderInfoList 分页获取tReceiveOrder表记录
func (tReceiveOrderService *TReceiveOrderService) GetTReceiveOrderInfoList(ctx context.Context, info systemReq.TReceiveOrderSearch) (list []system.TReceiveOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB_C.Model(&system.TReceiveOrder{})
	var tReceiveOrders []system.TReceiveOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tReceiveOrders).Error
	return tReceiveOrders, total, err
}
func (tReceiveOrderService *TReceiveOrderService) GetTReceiveOrderPublic(ctx context.Context) (list []system.TReceiveOrder, total int64, err error) {
	// 创建db
	db := global.GVA_DB_C.Model(&system.TReceiveOrder{}).Where("is_del = ?", "0")
	var tReceiveOrders []system.TReceiveOrder

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&tReceiveOrders).Error
	return tReceiveOrders, total, err
}
