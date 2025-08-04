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
	global.GVA_DB_Local = initialize.Gorm_Local() // gorm连接数据库
}

func TestTAddressActivitydService_CreateTAddressActivity(t *testing.T) {
	tAddressActivitydService := new(TAddressActivitydService)

	tAddressActivity := system.TAddressActivity{
		ActiveAddress:  "TJxE56ZJLitumu8DgsvpfdKcp1mJ3WAfjC",
		RegularAddress: "TPmsSsmjJGeMwHaQmxWCDtyUuVtbsvX1ZH",
		//AccountBalance:        3151.952,
		//CurrentOutboundAmount: 21.01,
		//AverageAmount:         21.111,
		//RecentThreeAvg:        9.212,
		OutboundCount:        12,
		CreatedBlock:         5777721212,
		CreatedTime:          time.Now(),
		LastTransactionBlock: 5777721212,
		LastTransactionTime:  time.Now(),
	}
	err := tAddressActivitydService.CreateTAddressActivity(context.Background(), &tAddressActivity)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tAddressActivity.Id)
}
func TestTAddressActivitydService_UpdateTAddressActivity(t *testing.T) {
	tAddressActivitydService := new(TAddressActivitydService)

	tAddressActivity := system.TAddressActivity{
		ActiveAddress:       "TJxE56ZJLitumu8DgsvpfdKcp1mJ3WAfjC",
		RegularAddress:      "TPmsSsmjJGeMwHaQmxWCDtyUuVtbsvX1ZH",
		TotalOutboundAmount: "999.999",

		LastTransactionTime: time.Now(),
		Id:                  1,
	}
	err := tAddressActivitydService.UpdateTAddressActivity(context.Background(), tAddressActivity)
	if err != nil {
		t.Error(err)
	}

}

func TestTAddressActivitydService_GetTAddressActivity(t *testing.T) {
	tAddressActivitydService := new(TAddressActivitydService)
	ActiveAddress := "TJxE56ZJLitumu8DgsvpfdKcp1mJ3WAfjC"
	RegularAddress := "TPmsSsmjJGeMwHaQmxWCDtyUuVtbsvX1ZH"
	record, err := tAddressActivitydService.GetTAddressActivity(context.Background(), ActiveAddress, RegularAddress)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(record.Id)
}
