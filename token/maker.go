package token

import "time"

type Maker interface {
	// CreateToken 根据用户名和时间创建一个新的token
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken 检查token是否有效
	VerifyToken(token string) (*Payload, error)
}
