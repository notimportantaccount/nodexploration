package nodexploration

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func uniswap() {
	// Connect to the Ethereum node
	client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
	if err != nil {
		fmt.Println(err)
		return
	}

	// USDC/ETH 0.05%
	// poolAddress := "0x990F5D39C304E2fc6Ece1D5242F32f317a21171D"

	blockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		return
	}

	fmt.Println(blockNum)
	// fmt.Println(common.HexToAddress(poolAddress))
	// // Get the total accumulated fees for the pool
	// fees, err := client.BalanceAt(context.Background(), common.HexToAddress(poolAddress), nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Total accumulated fees:", fees)
	// big.NewInt(2)

	// Get the list of addresses that have accumulated fees
	// var addresses []common.Address
	// for i := 0; i < int(fees.Int64()); i++ {
	// 	address, err := client.StorageAt(context.Background(), common.HexToAddress(poolAddress), big.NewInt(int64(i)), nil)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	addresses = append(addresses, address)
	// }
	// fmt.Println("Addresses that have accumulated fees:", addresses)

}
