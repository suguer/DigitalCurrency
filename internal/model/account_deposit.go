package model

import "time"

// AccountDeposit 用户在不同链上的余额信息
type AccountDeposit struct {
	ID        uint      `json:"id"`         // 记录ID
	AccountID uint      `json:"account_id"` // 关联的账户ID
	Chain     string    `json:"chain"`      // 链简称（如BTC、ETH等）
	Balance   float64   `json:"balance"`    // 余额数量
	Status    int       `json:"status"`     // 状态（0-禁用 1-正常）
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}
