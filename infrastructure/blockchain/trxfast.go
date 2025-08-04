package blockchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TronFastService struct {
	url      string
	username string
	password string
}

func NewTronFastService(url, username, password string) (*TronFastService, error) {
	return &TronFastService{url: url, username: username, password: password}, nil
}
func (s TronFastService) GetUserInfo() (*GetUserInfoResp, error) {
	// 创建请求体
	data := map[string]interface{}{
		"username": s.username,
		"password": s.password,
	}
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", s.url+"/api/getuserinfo", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result GetUserInfoResp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("JSON解析失败: %w", err)
	}
	return &result, nil
}

func (s TronFastService) BuyEnergy(re_address string, eneryAmount int64) (*BuyEnergyResp, error) {
	// 创建请求体
	data := map[string]interface{}{
		"username":   s.username,
		"password":   s.password,
		"re_type":    "ENERGY",    // 资源类型，当前仅支持'ENERGY'
		"re_address": re_address,  // 资源接收地址
		"re_value":   eneryAmount, // 资源数量
		"rent_time":  1,           // 租用时长，当前仅支持1，代表1小时
	}
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", s.url+"/api/buyenergy", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("购买失败：", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result BuyEnergyResp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("JSON解析失败: %w", err)
	}
	return &result, nil
}

func (s TronFastService) ActivateAddress(re_address string) (*ActivateAddressResp, error) {
	// 创建请求体
	data := map[string]interface{}{
		"username": s.username,
		"password": s.password,
		"address":  re_address, // 资源接收地址
	}
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", s.url+"/api/activateaddress", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("购买失败：", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("购买结果：", string(body))

	var result ActivateAddressResp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("JSON解析失败: %w", err)
	}
	return &result, nil
}

func (s TronFastService) OrderInfo(orderid string) (*OrderInfoResp, error) {
	// 创建请求体
	data := map[string]interface{}{
		"username": s.username,
		"password": s.password,
		"orderid":  orderid,
	}
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", s.url+"/api/orderinfo", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("购买结果：", string(body))

	var result OrderInfoResp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("JSON解析失败: %w", err)
	}
	return &result, nil
}

func (s TronFastService) AutoBuyEnergy(from_address, to_address string) (*AutoBuyEnergyResp, error) {
	// 创建请求体
	data := map[string]interface{}{
		"username":     s.username,
		"password":     s.password,
		"from_address": from_address, // 资源接收地址
		"to_address":   to_address,   // 资源接收地址
	}
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", s.url+"/api/autobuyenergy", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("购买失败：", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("购买结果：", string(body))

	var result AutoBuyEnergyResp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("JSON解析失败: %w", err)
	}
	return &result, nil
}
