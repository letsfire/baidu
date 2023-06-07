package baidu

import (
	"fmt"
	"time"
)

// TokenStorage TOKEN持久化
type TokenStorage interface {
	Store(key, value string, ttl time.Duration) error
	Fetch(key string) (string, error)
}

// localCache TOKEN本地缓存
type localCache struct {
	value     string
	expiredAt time.Time
}

func (tc *localCache) get() string {
	if tc.expiredAt.Before(time.Now()) {
		return ""
	}
	return tc.value
}

// tokenResponse TOKEN响应
type tokenResponse struct {
	ErrorCode   string `json:"error"`
	ErrorDesc   string `json:"error_description"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (res *tokenResponse) Error() error {
	if res.ErrorCode == "" {
		return nil
	}
	return fmt.Errorf("ai-sdk-baidu: fetch remote token error, code = %s, desc = %s", res.ErrorCode, res.ErrorDesc)
}
