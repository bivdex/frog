package system

import (
	"boost/data/server/core"
	"boost/data/server/global"
	"boost/data/server/initialize"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	global.GVA_VP = core.Viper("../../config.yaml")
	global.GVA_DB_B = initialize.Gorm_B() // gorm连接数据库
	// 初始化Viper
}
func TestGet24HourOrderStats(t *testing.T) {

	tOrderService := new(TOrderService)
	count, total, err := tOrderService.Get24HourOrderStats(context.Background(), "4567890")
	assert.Nil(t, err)

	fmt.Println(count, total, err)
	//assert.Equal(t, int64(1), count)
	//assert.Equal(t, 100.50, total)
}
func TestTReceiveOrderService_Get30MinOrderStats(t *testing.T) {
	tOrderService := new(TOrderService)
	count, err := tOrderService.Get30MinOrderStats(context.Background(), "4567890")
	assert.Nil(t, err)
	fmt.Println(count)
}
