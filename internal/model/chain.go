package model

import "time"

// Chain 区块链网络信息
type Chain struct {
	ID        uint      `json:"id"`         // 链ID
	Name      string    `json:"name"`       // 链名称
	Symbol    string    `json:"symbol"`     // 链符号
	Type      int       `json:"type"`       // 链类型（1-主链 2-测试链 3-私有链）
	Endpoint  string    `json:"endpoint"`   // 链节点地址
	Status    int       `json:"status"`     // 链状态（0-禁用 1-正常）
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}
