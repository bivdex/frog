package blockchain

import (
	"boost/data/server/infrastructure/blockchain_service"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"testing"
)

func TestBusinessService_QueryUSDTBalance(t *testing.T) {
	service, err := blockchain_service.NewTronService("https://nile.trongrid.io")
	//service, err := blockchain_service.NewTronService("https://api.trongrid.io")
	if err != nil {
		panic("Error while connecting to TRON RPC")
	}
	senderWallet := "TJxE56ZJLitumu8DgsvpfdKcp1mJ3WAfjC"
	receiverWallet := "TJxE56ZJLitumu8DgsvpfdKcp1mJ3WAfjC"
	tokenAddress := "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf"
	ctx := context.Background()

	businessService := NewBusinessService(service)
	senderBalance, _ := businessService.QueryUSDTBalance(ctx, senderWallet, receiverWallet, tokenAddress)

	fmt.Println("senderBalance:", senderBalance)
}
func TestQueryBalance(t *testing.T) {
	service, err := blockchain_service.NewTronService("https://nile.trongrid.io")
	//service, err := blockchain_service.NewTronService("https://api.trongrid.io")
	if err != nil {
		panic("Error while connecting to TRON RPC")
	}
	senderWallet := "TJxE56ZJLitumu8DgsvpfdKcp1mJ3WAfjC"
	receiverWallet := "TPmsSsmjJGeMwHaQmxWCDtyUuVtbsvX1ZH"
	ctx := context.Background()

	businessService := NewBusinessService(service)
	senderBalance, receiverBalance := businessService.QueryBalance(ctx, senderWallet, receiverWallet)

	fmt.Println("senderBalance:", senderBalance)
	fmt.Println("receiverBalance:", receiverBalance)

}
func TestTransferNative(t *testing.T) {
	service, err := blockchain_service.NewTronService("https://nile.trongrid.io")
	//service, err := blockchain_service.NewTronService("https://api.trongrid.io")
	if err != nil {
		panic("Error while connecting to TRON RPC")
	}
	//senderWallet := "TJxE56ZJLitumu8DgsvpfdKcp1mJ3WAfjC"
	receiverWallet := "TPmsSsmjJGeMwHaQmxWCDtyUuVtbsvX1ZH"
	ctx := context.Background()
	businessService := NewBusinessService(service)
	hash, err := businessService.Transfer(ctx, receiverWallet, 42.011, "43b8e682fd65cfc5fd0a67d0caf6c5451e271aacb2f055d6c5f2c429470e0e23")

	if err != nil {
		fmt.Errorf("transfer failure")
	}
	fmt.Println("hash:", hash)
}
func TestTransferToken(t *testing.T) {
	service, err := blockchain_service.NewTronService("https://nile.trongrid.io")
	//service, err := blockchain_service.NewTronService("https://api.trongrid.io")
	if err != nil {
		panic("Error while connecting to TRON RPC")
	}
	//senderWallet := "TJxE56ZJLitumu8DgsvpfdKcp1mJ3WAfjC"
	receiverWallet := "TPmsSsmjJGeMwHaQmxWCDtyUuVtbsvX1ZH"
	ctx := context.Background()
	businessService := NewBusinessService(service)
	hash, err := businessService.TransferToken(ctx, receiverWallet, "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf", 42.011, "43b8e682fd65cfc5fd0a67d0caf6c5451e271aacb2f055d6c5f2c429470e0e23")

	if err != nil {
		fmt.Errorf("transfer failure")
	}
	fmt.Println("hash:", hash)
}

func TestBusinessService_GetAddressFromSecretKey(t *testing.T) {

	service, err := blockchain_service.NewTronService("https://nile.trongrid.io")
	//service, err := blockchain_service.NewTronService("https://api.trongrid.io")
	if err != nil {
		panic("Error while connecting to TRON RPC")
	}
	businessService := NewBusinessService(service)
	fromAddress, err := businessService.GetAddressFromSecretKey("43b8e682fd65cfc5fd0a67d0caf6c5451e271aacb2f055d6c5f2c429470e0e23")

	if err != nil {
		fmt.Errorf("get address failure")
	}
	log.Println("fromAddress:", fromAddress)
}
