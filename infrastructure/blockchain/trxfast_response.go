package blockchain

type ActivateAddressResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Address string  `json:"address"`
		Price   float64 `json:"price"`
		Day     string  `json:"day"`
	} `json:"data"`
}

type AutoBuyEnergyResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderID    int64   `json:"orderId"`
		Type       string  `json:"type"`
		RentTime   int     `json:"rent_time"`
		Value      int     `json:"value"`
		OrderPrice float64 `json:"orderPrice"`
		ToAddress  string  `json:"toAddress"`
		Hash       string  `json:"hash"`
		CreateTime string  `json:"createTime"`
	} `json:"data"`
}
type OrderInfoResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Orderid   int64   `json:"orderid"`
		ReAddress string  `json:"re_address"`
		ReValue   int     `json:"re_value"`
		RentTime  int     `json:"rent_time"`
		Price     float64 `json:"price"`
		Hash      string  `json:"hash"`
		Status    int     `json:"status"`
		Ordertime string  `json:"ordertime"`
	} `json:"data"`
}
type BuyEnergyResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderID    int64   `json:"orderId"`
		Type       string  `json:"type"`
		RentTime   int     `json:"rent_time"`
		Value      int     `json:"value"`
		OrderPrice float64 `json:"orderPrice"`
		ToAddress  string  `json:"toAddress"`
		Hash       string  `json:"hash"`
		CreateTime string  `json:"createTime"`
	} `json:"data"`
}
type GetUserInfoResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Username        string `json:"username"`
		Balance         int    `json:"balance"`
		RechargeAddress string `json:"rechargeAddress"`
		TotalTecharge   int    `json:"totalTecharge"`
		Amountused      int    `json:"amountused"`
		UsageCount      int    `json:"usageCount"`
		Status          int    `json:"status"`
	} `json:"data"`
}
