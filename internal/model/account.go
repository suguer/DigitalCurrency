package model

import "time"

// Account 用户账户信息
type Account struct {
	ID        uint      `json:"id"`         // 账户ID
	Username  string    `json:"username"`   // 用户名
	Password  string    `json:"password"`   // 密码（加密存储）
	Email     string    `json:"email"`      // 邮箱
	Phone     string    `json:"phone"`      // 手机号
	Status    int       `json:"status"`     // 账户状态（0-禁用 1-正常）
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	Address   string    `json:"address"`    // 钱包地址
}
