package blockchain_service

import (
	"fmt"
	"log"
	"math/big"
	"testing"
)

func TestEstimateEnrgyandBandwidth(t *testing.T) {

	tronService, _ := NewTronService("https://young-clean-orb.tron-mainnet.quiknode.pro/9283c9ddb51102d9d22cf1ac5a6e6fc898eeaf77")

	// 示例参数
	fromAddress := "THjixxAv9hCSC25WVqxBdBSDDP96hvRGCe" // 替换为你的地址
	toAddress := "TPXH2iHQY6V58uPy8LYAc2gqqyGZtXpKBN"   // 替换为接收地址
	amount := big.NewInt(1_000_000_000_000)             // 100 USDT (6 decimals)

	// 估算USDT转账费用
	energy, bandwidth, fee, err := tronService.EstimateUSDTTransfer(fromAddress, toAddress, amount)
	if err != nil {
		log.Fatalf("估算失败: %v", err)
	}

	// 打印结果
	fmt.Printf("\nUSDT转账资源估算:\n")
	fmt.Printf("├─ 能量消耗: %d Energy\n", energy)
	fmt.Printf("├─ 带宽消耗: %d Bandwidth\n", bandwidth)
	fmt.Printf("├─ 预估手续费: %.6f TRX\n", fee)
	fmt.Printf("└─ 备注: 实际费用取决于网络状况\n\n")
}
