package system

import (
	"boost/data/server/core"
	"boost/data/server/global"
	"boost/data/server/initialize"
	"boost/data/server/model/system"
	"context"
	"log"
	"testing"
)

func init() {
	global.GVA_VP = core.Viper("../../config.yaml")
	global.GVA_DB_D = initialize.Gorm_D() // gorm连接数据库
	global.GVA_DB = initialize.Gorm_A()   // gorm连接数据库
	global.GVA_DB_B = initialize.Gorm_B() // gorm连接数据库
	// 初始化Viper
}

func TestTOrderToAddressRecordService_CreateTOrderToAddressRecord(t *testing.T) {

	tOrderAddressRecordResultService := new(TOrderAddressRecordResultService)

	result, _ := tOrderAddressRecordResultService.GetTOrderAddressRecordResultByFromAddressPart(context.Background(), "1234567")

	log.Println(result)
	tOrderToAddressRecordService := new(TOrderToAddressRecordService)
	tOrder := system.TOrderToAddressRecord{
		FromAddressPart: result.FromAddressPart,
		ToAddress:       result.Address,
		CreateTime:      result.MatchSuccessTime,
	}
	err := tOrderToAddressRecordService.CreateTOrderToAddressRecord(context.Background(), &tOrder)

	if err != nil {
		t.Error(err)
	}

}
