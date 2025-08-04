package blockchain

import (
	"boost/data/server/infrastructure/blockchain_service"
	"boost/data/server/pkg/tron"
	"boost/data/server/utils"
	"context"
	"fmt"
)

type BusinessService struct {
	service *blockchain_service.TronService
}

func NewBusinessService(service *blockchain_service.TronService) *BusinessService {
	return &BusinessService{service: service}
}

// CheckEnergyBalance 检查账户能量是否足够执行操作
func (business BusinessService) GetEnergyBalance(address string) (int64, int64, error) {
	energy, bandwidth, err := business.service.GetEnergyBalance(address)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get account: %w", err)
	}

	return energy, bandwidth, nil
}

func (business BusinessService) IsAcccountActive(ctx context.Context, sender string) bool {
	_, err := business.service.IsAcccountActive(ctx, sender)
	if err != nil {
		return false
	}
	return true
}

func (business BusinessService) EstimateUSDTTransfer(fromAddr, toAddr string, _amount float64) (int64, int64, float64, error) {
	sendAmount := utils.ConvertFloatToBigInt(_amount, 6)

	return business.service.EstimateUSDTTransfer(fromAddr, toAddr, sendAmount)
}

func (business BusinessService) GetAddressFromSecretKey(privateKey string) (string, error) {
	return tron.GetTronAddressFromPrivateKey(privateKey)
}

func (business BusinessService) QueryUSDTBalanceForAddress(
	ctx context.Context,
	sender string, tokenAddress string) float64 {

	senderWalletBalance, err := business.service.GetTokenBalance(ctx, sender, tokenAddress)
	if err != nil {
		return 0
	}
	senderWalletBalanceFloat := utils.ConvertBigIntToFloat(
		senderWalletBalance, 6,
	)
	return senderWalletBalanceFloat
}
func (business BusinessService) QueryUSDTBalance(
	ctx context.Context,
	sender string, receiver string, tokenAddress string) (float64, float64) {
	////判断地址是否激活状态
	//avaiable, _ := business.service.IsAcccountActive(ctx, sender)
	//
	//if !avaiable {
	//	//主地址进行转移2trx，进行激活
	//	_, err := business.Transfer(ctx, sender, 2, "")
	//	if err != nil {
	//		return 0, 0
	//	}
	//}

	senderWalletBalance, err := business.service.GetTokenBalance(ctx, sender, tokenAddress)
	if err != nil {
		return 0, 0
	}
	senderWalletBalanceFloat := utils.ConvertBigIntToFloat(
		senderWalletBalance, 6,
	)

	receiverWalletBalance, err := business.service.GetTokenBalance(ctx, receiver, tokenAddress)
	if err != nil {
		return 0, 0
	}
	receiverWalletBalanceFloat := utils.ConvertBigIntToFloat(
		receiverWalletBalance, 6,
	)

	//receiverWalletBalanceFloat>0 ，执行转账前，需先从能量接口调用能量，对方地址有U转65000能量，对方地址无U转130000能量

	return senderWalletBalanceFloat, receiverWalletBalanceFloat
}
func (business BusinessService) QueryBalanceForAddress(
	ctx context.Context,
	sender string,
) float64 {
	senderWalletBalance, err := business.service.GetNativeBalance(ctx, sender)
	if err != nil {
		panic(err)
	}
	senderWalletBalanceFloat := utils.ConvertBigIntToFloat(
		senderWalletBalance, 6,
	)

	return senderWalletBalanceFloat
}

func (business BusinessService) QueryBalance(
	ctx context.Context,
	sender,
	receiver string,
) (float64, float64) {
	////判断地址是否激活状态
	//avaiable, _ := business.service.IsAcccountActive(ctx, sender)
	//
	//if !avaiable {
	//	//主地址进行转移2trx，进行激活
	//	_, err := business.Transfer(ctx, sender, 2, "")
	//	if err != nil {
	//		return 0, 0
	//	}
	//}

	senderWalletBalance, err := business.service.GetNativeBalance(ctx, sender)
	if err != nil {
		panic(err)
	}
	receiverWalletBalance, err := business.service.GetNativeBalance(ctx, receiver)
	if err != nil {
		panic(err)
	}
	senderWalletBalanceFloat := utils.ConvertBigIntToFloat(
		senderWalletBalance, 6,
	)
	receiverWalletBalanceFloat := utils.ConvertBigIntToFloat(
		receiverWalletBalance, 6,
	)
	fmt.Printf(
		"Sender wallet balance: %f TRX\nReceiver wallet balance: %f TRX",
		senderWalletBalanceFloat,
		receiverWalletBalanceFloat,
	)
	return senderWalletBalanceFloat, receiverWalletBalanceFloat
}

func (business BusinessService) Transfer(ctx context.Context, receiverWallet string, _amount float64, senderPk string) (string, error) {
	sendAmount := utils.ConvertFloatToBigInt(_amount, 6)
	result, err := business.service.TransferNative(
		ctx,
		senderPk,
		receiverWallet,
		sendAmount,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return result.String(), err
}

func (business BusinessService) TransferAll(ctx context.Context, receiverWallet string, senderPk string) {
	result, err := business.service.TransferNativeRemaining(
		ctx,
		senderPk,
		receiverWallet,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func (business BusinessService) TransferAllToken(ctx context.Context, receiverWallet string, senderPk string) {
	tokenAddress := "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf"
	res, err := business.service.TransferTokenRemaining(
		ctx,
		senderPk,
		tokenAddress,
		receiverWallet,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func (business BusinessService) TransferToken(ctx context.Context, receiverWallet string, _contract string, _amount float64, senderPk string) (string, error) {
	tokenAddress := _contract
	sendAmount := utils.ConvertFloatToBigInt(_amount, 6)
	res, err := business.service.TransferToken(
		ctx,
		senderPk,
		tokenAddress,
		receiverWallet,
		sendAmount,
	)
	if err != nil {
		panic(err)
	}
	return res.String(), err
}
