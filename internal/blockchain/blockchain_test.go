package blockchain

import (
	"DigitalCurrency/internal/config"
	"context"
	"fmt"
	"testing"
)

func TestTronGetNowBlockId(t *testing.T) {
	conf := &config.EthConfig{
		Chain: "tron",
		Node:  "https://api.trongrid.io",
		//ApiKey:
	}
	tron := NewTron(context.Background(), conf)
	BlockId, Time, err := tron.GetNowBlockId()
	fmt.Printf("BlockId: %v\n", BlockId)
	fmt.Printf("Time: %v\n", Time)
	fmt.Printf("err: %v\n", err)
}
func TestTronGetBlockByNum(t *testing.T) {
	conf := &config.EthConfig{
		Chain: "tron",
		Node:  "https://api.trongrid.io",
	}
	tron := NewTron(context.Background(), conf)
	data, err := tron.GetBlockByNum(71229398)
	fmt.Printf("BlockID: %v\n", data.BlockID)
	fmt.Printf("BlockHeader: %v\n", data.BlockHeader)
	fmt.Printf("Transactions: %+v\n", data.Transactions[0])
	fmt.Printf("err: %v\n", err)
}
