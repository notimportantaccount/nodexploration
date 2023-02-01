package nodexploration

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func aaveDepositEvents() {
	// Connect to the Ethereum node
	client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
	if err != nil {
		fmt.Println(err)
		return
	}

	blockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		return
	}

	fmt.Println(blockNum)

	q := ethereum.FilterQuery{
		FromBlock: big.NewInt(16530480),
		ToBlock:   big.NewInt(int64(blockNum)),
		Topics:    [][]common.Hash{[]common.Hash{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
	}

	logs, err := client.FilterLogs(context.Background(), q)
	if err != nil {
		return
	}

	fmt.Println(logs[0])
}
