// 自动生成模板TReceiveOrder
package system

// tReceiveOrder表 结构体  TReceiveOrder
type TReceiveOrder struct {
	//ID              uint    `gorm:"primarykey" json:"ID"`                                                                           // 主键ID                         //id字段
	Cztimes         int     `json:"cz_times" form:"cz_times" gorm:"column:cz_times;comment:cz_times;"`                              //cztimes
	OrderNo         string  `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;size:100;"`                            //订单号
	FromAddressPart string  `json:"fromAddressPart" form:"fromAddressPart" gorm:"column:from_address_part;comment:前3后4码;size:255;"` //前3后4码
	ToAddress       string  `json:"toAddress" form:"toAddress" gorm:"column:to_address;comment:目标地址;size:255;"`                     //目标地址
	Amount          float64 `json:"amount" form:"amount" gorm:"column:amount;comment:;size:64;"`                                    //amount字段
	OrderTime       string  `json:"orderTime" form:"orderTime" gorm:"column:order_time;comment:订单时间;size:255;"`                     //订单时间
	CreateTime      string  `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                           //创建时间
	//Initialization  bool      `json:"initialization" form:"initialization" gorm:"column:initialization;comment:是否初始化:0未初始化 1.已初始化;"`  //是否初始化:0未初始化 1.已初始化
	//ErrorData       bool      `json:"errorData" form:"errorData" gorm:"column:error_data;comment:异常数据:1正常.2.重复 3.金额大于1000;"`          //异常数据:1正常.2.重复 3.金额大于1000
	//WaitMatch       bool      `json:"waitMatch" form:"waitMatch" gorm:"column:wait_match;comment:是否匹配:0.等待匹配 1.已匹配;"`                 //是否匹配:0.等待匹配 1.已匹配
	IsDel string `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否删除 0无，1已删;"` //是否删除 0无，1已删
}

// TableName tReceiveOrder表 TReceiveOrder自定义表名 t_receive_order
func (TReceiveOrder) TableName() string {
	return "t_receive_order"
}
