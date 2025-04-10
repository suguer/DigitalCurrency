package model

import "time"

// Transaction 区块链交易信息
type Transaction struct {
	ID            uint      `json:"id"`            // 交易ID
	Hash          string    `json:"hash"`          // 交易哈希
	Chain         string    `json:"chain"`         // 链简称
	FromAddress   string    `json:"from_address"`  // 发送方地址
	ToAddress     string    `json:"to_address"`    // 接收方地址
	Amount        float64   `json:"amount"`        // 交易金额
	Fee           float64   `json:"fee"`           // 交易手续费
	Status        int       `json:"status"`        // 交易状态（0-待确认 1-已确认 2-失败）
	BlockHash     string    `json:"block_hash"`    // 所属区块哈希
	BlockHeight   uint64    `json:"block_height"`  // 所属区块高度
	Confirmations uint64    `json:"confirmations"` // 确认数
	CreatedAt     time.Time `json:"created_at"`    // 创建时间
	ConfirmedAt   time.Time `json:"confirmed_at"`  // 确认时间
}
