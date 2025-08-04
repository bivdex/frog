package system

import (
	"boost/data/server/global"
	"boost/data/server/utils"
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
)

//func init() {
//	global.GVA_VP = core.Viper("../../config.yaml")
//	global.GVA_DB_C = initialize.Gorm_C() // gorm连接数据库
//}

func TestTReceiveOrderService_CreateTReceiveOrder(t *testing.T) {

	//currentTime := time.Now()
	//tReceiveOrderService := new(TReceiveOrderService)
	formattedAmount := ""
	numStr := "2222.3333"

	integerPart := utils.GetIntegerPart(numStr)
	fa := utils.SuoJinSuanFa(integerPart)
	switch global.SUOJINSUANFA {
	case 1:
		formattedAmount = fa
	case 2:
		formattedAmount = utils.SuoJinSuanFa2(integerPart)
	case 3:
		formattedAmount = utils.SuoJinSuanFa3(integerPart)
	default:
		formattedAmount = fa
	}

	log.Println("缩进算法金额前", numStr)
	log.Println("缩进算法金额后", formattedAmount)
	_amount, _ := strconv.ParseFloat(formattedAmount, 64)

	log.Println(_amount)
	//treceiveOrder := system.TReceiveOrder{
	//	//Id:    "23803",
	//	IsDel:           "1",
	//	Cztimes:         1,
	//	OrderNo:         "1",
	//	FromAddressPart: "1",
	//	ToAddress:       "1",
	//	Amount:          _amount,
	//	OrderTime:       "202509012",
	//	CreateTime:      currentTime.Format(time.DateTime),
	//}
	//err := tReceiveOrderService.CreateTReceiveOrder(context.Background(), &treceiveOrder)
	//assert.Nil(t, err)
}

//func TestTReceiveOrderServiceUpdateTReceiveOrder(t *testing.T) {
//	tReceiveOrderService := new(TReceiveOrderService)
//	treceiveOrder := system.TReceiveOrder{
//		//Id:    "23803",
//		IsDel: "1",
//	}
//	err := tReceiveOrderService.UpdateTReceiveOrder(context.Background(), treceiveOrder)
//	assert.Nil(t, err)
//}

func TestTReceiveOrderService_GetTReceiveOrderPublic(t *testing.T) {
	tReceiveOrderService := new(TReceiveOrderService)

	list, total, err := tReceiveOrderService.GetTReceiveOrderPublic(context.Background())

	if err != nil {
	}

	fmt.Println(list)
	fmt.Println(total)
}
func TestTReceiveOrderService_GetTReceiveOrderByOrderNo(t *testing.T) {
	tReceiveOrderService := new(TReceiveOrderService)

	record, err := tReceiveOrderService.GetTReceiveOrderByOrderNo(context.Background(), "202505100926NVBE")

	if err != nil {
		fmt.Println("错误： ", err)
	} else {
		fmt.Println("正确： ", err)
		fmt.Println(record)

	}

}
func TestTReceiveOrderService_GetTReceiveOrderByFromAddressPart(t *testing.T) {
	tReceiveOrderService := new(TReceiveOrderService)

	records, total, err := tReceiveOrderService.GetTReceiveOrderByFromAddressPart(context.Background(), "202505100926NVBE")

	if err != nil {
		fmt.Println("错误： ", err)
	} else {
		fmt.Println("正确： ", err)
		fmt.Println(total)
		fmt.Println(records)
	}

}
