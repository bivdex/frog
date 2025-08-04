package system

import (
	"boost/data/server/global"
	"boost/data/server/model/system"
	systemReq "boost/data/server/model/system/request"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type TOrderService struct{}

// CreateTOrder 创建tOrder表记录
func (tOrderService *TOrderService) CreateTOrder(ctx context.Context, tOrder *system.TOrder) (err error) {
	err = global.GVA_DB_B.Create(tOrder).Error
	return err
}

// DeleteTOrder 删除tOrder表记录
func (tOrderService *TOrderService) DeleteTOrder(ctx context.Context, id string) (err error) {
	err = global.GVA_DB_B.Delete(&system.TOrder{}, "id = ?", id).Error
	return err
}

// DeleteTOrderByIds 批量删除tOrder表记录
func (tOrderService *TOrderService) DeleteTOrderByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB_B.Delete(&[]system.TOrder{}, "id in ?", ids).Error
	return err
}

// UpdateTOrder 更新tOrder表记录
func (tOrderService *TOrderService) UpdateTOrder(ctx context.Context, tOrder system.TOrder) (err error) {
	err = global.GVA_DB_B.Model(&system.TOrder{}).Where("id = ?", tOrder.Id).Updates(&tOrder).Error
	return err
}

// GetTOrder 根据id获取tOrder表记录
func (tOrderService *TOrderService) GetTOrder(ctx context.Context, orderNO string) (tOrder system.TOrder, err error) {
	err = global.GVA_DB_B.Where("order_no = ?", orderNO).First(&tOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tOrder, nil
	} else if err != nil {
		return tOrder, err
	}
	return
}

// GetTOrder 根据id获取tOrder表记录
func (tOrderService *TOrderService) GetTOrderByOrderNO(ctx context.Context, orderNO string) (tOrder system.TOrder, err error) {
	err = global.GVA_DB_B.Where("order_no = ?", orderNO).First(&tOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tOrder, nil
	} else if err != nil {
		return tOrder, err
	}
	return
}

// GetTOrder 根据id获取tOrder表记录
func (tOrderService *TOrderService) GetTOrderByAddress(ctx context.Context, from_address string) (tOrder system.TOrder, err error) {
	err = global.GVA_DB_B.Where("from_address = ?", from_address).First(&tOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return tOrder, nil
	} else if err != nil {
		return tOrder, err
	}
	return
}

// GetTOrderInfoList 分页获取tOrder表记录
func (tOrderService *TOrderService) GetTOrderInfoList(ctx context.Context, info systemReq.TOrderSearch) (list []system.TOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.TOrder{})
	var tOrders []system.TOrder
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
func (tOrderService *TOrderService) GetTOrderPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

/**
 * Counts orders created in the last 24 hours 统计最近24小时内的某个地址的订单数量
 * @return number of orders
 */
func (tOrderService *TOrderService) CountOrdersLast24Hours(ctx context.Context, _fromAddress string) (int64, error) {
	// 此方法为获取数据源定义的数据
	// 请自行实现

	return 0, nil
}
func (tOrderService *TOrderService) SumOrderAmountLast24Hours(ctx context.Context, _fromAddress string) (int64, error) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
	return 0, nil
}

/**
 * Counts orders created in the last 24 hours
 * @return number of orders
 */
func (tOrderService *TOrderService) Get24HourOrderStats(ctx context.Context, _fromAddressPart string) (count int64, totalAmount float64, err error) {
	// 计算24小时前的时间点
	startTime := time.Now().Add(-24 * time.Hour)
	db := global.GVA_DB_B.Model(&system.TOrder{})
	// 执行查询
	err = db.
		Where("transfer_time >= ?", startTime).
		Where("from_address_part = ?", _fromAddressPart).
		//Where("status = ?", 9).
		Select("COUNT(*) as count, COALESCE(SUM(amount), 0) as total").
		Row().
		Scan(&count, &totalAmount)
	return
}
func (tOrderService *TOrderService) Get30MinOrderStats(ctx context.Context, fromAddressPart string) (count int64, err error) {
	startTime := time.Now().Add(-30 * time.Minute)
	db := global.GVA_DB_B.Model(&system.TOrder{})
	err = db.Where("transfer_time >= ?", startTime).
		Where("from_address_part = ?", fromAddressPart).
		Count(&count).Error
	return count, err
}
