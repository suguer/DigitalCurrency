package mdb

import "time"

// Wallet 钱包信息
type Wallet struct {
	ID         uint      `json:"id"`          // 钱包ID
	Address    string    `json:"address"`     // 钱包地址
	PrivateKey string    `json:"private_key"` // 密钥
	Status     int       `json:"status"`      // 钱包状态（0-禁用 1-正常）
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
	UpdatedAt  time.Time `json:"updated_at"`  // 更新时间
}
