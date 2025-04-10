package model

import "time"

// Order 订单信息
type Order struct {
	ID            uint      `json:"id"`             // 订单ID
	OrderNo       string    `json:"order_no"`       // 订单号
	Chain         string    `json:"chain"`          // 链简称
	AccountID     uint      `json:"account_id"`     // 关联的账户ID
	TransactionID uint      `json:"transaction_id"` // 关联的交易ID
	Type          int       `json:"type"`           // 订单类型（1-充值 2-提现 3-转账）
	Amount        float64   `json:"amount"`         // 订单金额
	Fee           float64   `json:"fee"`            // 手续费
	Status        int       `json:"status"`         // 订单状态（0-待处理 1-处理中 2-已完成 3-已取消 4-失败）
	Remark        string    `json:"remark"`         // 备注
	CreatedAt     time.Time `json:"created_at"`     // 创建时间
	UpdatedAt     time.Time `json:"updated_at"`     // 更新时间
	CompletedAt   time.Time `json:"completed_at"`   // 完成时间
}
