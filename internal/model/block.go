package model

import "time"

// Block 区块信息
type Block struct {
	Height    uint64    `json:"height"`    // 区块高度
	Hash      string    `json:"hash"`      // 区块哈希
	PrevHash  string    `json:"prev_hash"` // 前一个区块的哈希
	Timestamp time.Time `json:"timestamp"` // 区块时间戳
}
