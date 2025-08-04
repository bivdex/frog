package system

import (
	"boost/data/server/core"
	"boost/data/server/global"
	"boost/data/server/initialize"
	"boost/data/server/model/common/request"
	systemReq "boost/data/server/model/system/request"
	"context"
	"fmt"
	"log"
	"testing"
)

func init() {
	global.GVA_VP = core.Viper("../../config.yaml")
	global.GVA_DB_E = initialize.Gorm_E() // gorm连接数据库
	// 初始化Viper
}

func TestTOrderTransferFromAddressService_CreateTOrderTransferFromAddress(t *testing.T) {
	tOrderService := new(TOrderFromAddressService)
	record, err := tOrderService.GetTOrderFromAddress(context.Background(), "1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(record)
}

func TestTOrderTransferFromAddressService_GetTOrderInfoList(t *testing.T) {
	tOrderService := new(TOrderFromAddressService)
	pageInfo := request.PageInfo{
		Page:     2,
		PageSize: 100,
	}
	info := systemReq.TOrderSearch{
		PageInfo: pageInfo,
	}

	list, total, _ := tOrderService.GetTOrderInfoList(context.Background(), info)

	log.Println(total)
	for _, item := range list {
		fmt.Println(item.FromAddress)
	}
}
func TestTOrderTransferFromAddressService_GetFullTOrderInfoList(t *testing.T) {
	tOrderService := new(TOrderFromAddressService)

	list, total, _ := tOrderService.GetFullTOrderInfoList(context.Background())

	log.Println(total)
	for _, item := range list {
		fmt.Println(item.FromAddress)
	}
}
