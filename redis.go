package baidu

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/letsfire/redigo/v2"
)

type redisStorage struct {
	client *redigo.Client
}

func (rs redisStorage) Fetch(key string) (string, error) {
	v, err := rs.client.String(func(c redis.Conn) (interface{}, error) {
		return c.Do("GET", key)
	})
	if err == redis.ErrNil {
		return "", nil
	}
	return v, err
}

func (rs redisStorage) Store(key, value string, ttl time.Duration) error {
	_, err := rs.client.Execute(func(c redis.Conn) (interface{}, error) {
		return c.Do("PSETEX", key, ttl.Milliseconds(), value)
	})
	return err
}

func RedisStorage(client *redigo.Client) TokenStorage {
	return redisStorage{client: client}
}
