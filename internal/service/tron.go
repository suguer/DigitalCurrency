package service

import (
	"DigitalCurrency/internal/blockchain"
	"DigitalCurrency/internal/config"
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TronService struct {
	client      *blockchain.Tron
	LastBlockId int
	chBlock     chan int
	workerCount int
	conf        *config.EthConfig
}

var (
	tron     TronService
	tronOnce sync.Once
)

func NewTronService(ctx context.Context, conf *config.EthConfig) *TronService {
	tronOnce.Do(func() {
		tron = TronService{
			client:      blockchain.NewTron(ctx, conf),
			chBlock:     make(chan int, 20),
			workerCount: 5,
			conf:        conf,
		}
	})
	return &tron
}

func (t *TronService) Start() {
	for i := 0; i < t.workerCount; i++ {
		go t.Handler()
	}
	for {
		BlockId, _, err := t.client.GetNowBlockId()
		fmt.Printf("BlockId: %v\n", BlockId)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			time.Sleep(10 * time.Second)
			continue
		}
		if t.LastBlockId == 0 {
			t.LastBlockId = BlockId - 10
		}
		for i := t.LastBlockId; i < BlockId; i++ {
			t.chBlock <- i
		}
		t.LastBlockId = BlockId
		time.Sleep(10 * time.Second)
	}

}

func (t *TronService) Handler() {
	ContractMap := t.conf.ContractMap()
	for blocksnum := range t.chBlock {
		fmt.Printf("blocksnum: %v\n", blocksnum)
		data, err := t.client.GetBlockByNum(blocksnum)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			if strings.Contains(err.Error(), "unexpected end of JSON input") {
				t.chBlock <- blocksnum
			}
			continue

		}
		for _, transaction := range data.Transactions {
			out_trade_no := transaction.TxID
			status := strings.ToLower(transaction.Ret[0].ContractRet)
			if status != "success" {
				continue
			}
			contractAddress := transaction.RawData.Contract[0].Parameter.Value.ContractAddress
			if _, ok := ContractMap[contractAddress]; !ok {
				continue
			}
			funcValue := transaction.RawData.Contract[0].Parameter.Value.Data[0:8]
			if funcValue != "a9059cbb" {
				continue
			}
			toAddress := strings.TrimLeft(transaction.RawData.Contract[0].Parameter.Value.Data[9:72], "0")

			amountHex := strings.TrimLeft(transaction.RawData.Contract[0].Parameter.Value.Data[72:], "0")
			if amountHex == "" {
				amountHex = "0"
			}
			amountInt, err := strconv.ParseInt(amountHex, 16, 64)
			if err != nil {
				fmt.Printf("Error parsing amount: %v\n", err)
				continue
			}
			transactionAmount := float64(amountInt) / 1000000

			fmt.Printf("blocksnum:%v,out_trade_no: %v,contact:%v,to:%v,amount:%v\n", blocksnum, out_trade_no, contractAddress, toAddress, transactionAmount)
			// TODO: 判断是否需要激活该笔转账记录
		}
	}
}
