package system

import (
	"boost/data/server/core"
	"boost/data/server/global"
	"boost/data/server/initialize"
	"boost/data/server/model/system"
	"context"
	"fmt"
	"testing"
	"time"
)

func init() {
	global.GVA_VP = core.Viper("../../config.yaml")
	global.GVA_DB = initialize.Gorm_A() // gorm连接数据库
	// 初始化Viper
}

func TestTOrderAddressRecordResultService_GetTOrderAddressRecordResultByFromAddressPart(t *testing.T) {
	tOrderAddressRecordResultService := new(TOrderAddressRecordResultService)
	record, err := tOrderAddressRecordResultService.GetTOrderAddressRecordResultByFromAddressPart(context.Background(), "11234567")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(record.Id)
	//minutes := record.MatchSuccessTime.Sub(record.CreateTime).Abs().Minutes()
	//fmt.Println(record)
	//fmt.Println(minutes)
}
func TestTOrderAddressRecordResultService_CreateTOrderAddressRecordResult(t *testing.T) {
	tOrderAddressRecordResultService := new(TOrderAddressRecordResultService)
	tOrderAddressRecordResult := system.TOrderAddressRecordResult{
		FromAddressPart:  "1234567",
		Address:          "TPmsSsmjJGeMwHaQmxWCDtyUuVtbsvX1ZH",
		PrivateAddress:   "43b8e682fd65cfc5fd0a67d0caf6c5451e271aacb2f055d6c5f2c429470e0e23",
		CreateTime:       time.Now(),
		MatchSuccessTime: time.Now()}
	err := tOrderAddressRecordResultService.CreateTOrderAddressRecordResult(context.Background(), &tOrderAddressRecordResult)
	if err != nil {
		t.Error(err)
	}

}
