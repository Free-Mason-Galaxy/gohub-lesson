// Package captcha
// descr 配置验证码
// author fm
// date 2022/11/16 16:38
package captcha

import (
	"errors"
	"time"

	"gohub-lesson/pkg/app"
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/redis"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (class *RedisStore) Set(key string, value string) error {

	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))

	// 方便本地开发调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}

	if ok := class.RedisClient.Set(class.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}

	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (class *RedisStore) Get(key string, clear bool) string {
	key = class.KeyPrefix + key
	val := class.RedisClient.Get(key)
	if clear {
		class.RedisClient.Del(key)
	}
	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (class *RedisStore) Verify(key, answer string, clear bool) bool {
	v := class.Get(key, clear)
	return v == answer
}
