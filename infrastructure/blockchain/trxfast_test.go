package blockchain

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestTronFastService_ActivateAddress(t *testing.T) {
	baseURL := "https://trxfast.com"
	trxfast, err := NewTronFastService(baseURL, "testuser", "testuser")
	if err != nil {
		fmt.Errorf("invalid username or password")
	}
	re_address := "TPmsSsmjJGeMwHaQmxWCDtyUuVtbsvX1ZH"
	resp, err := trxfast.ActivateAddress(re_address)
	if err != nil {
		fmt.Errorf("no reason")
	}
	fmt.Println(resp)
}

func TestTronFastService_GetUserInfo(t *testing.T) {
	baseURL := "https://trxfast.com"
	trxfast, err := NewTronFastService(baseURL, "testuser", "testuser")
	if err != nil {
		fmt.Errorf("invalid username or password")
	}
	resp, err := trxfast.GetUserInfo()
	if err != nil {
		fmt.Errorf("no reason")
	}
	fmt.Println(resp)
	assert.Equal(t, resp.Code, 10000)
}
func TestTronFastService_BuyEnergy(t *testing.T) {
	baseURL := "https://trxfast.com"
	trxfast, err := NewTronFastService(baseURL, "testuser", "testuser")
	if err != nil {
		fmt.Errorf("invalid username or password")
	}
	resp, err := trxfast.GetUserInfo()
	if err != nil {
		fmt.Errorf("no reason")
	}
	fmt.Println(resp)
	assert.Equal(t, resp.Code, 10000)
}
