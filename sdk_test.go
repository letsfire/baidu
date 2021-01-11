package baidu

import (
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/letsfire/redigo/v2/mode/alone"
	"github.com/stretchr/testify/assert"
)

var rds = alone.NewClient(
	alone.Addr("redis-01.data.svc.cluster.local:30003"),
	alone.DialOpts(
		redis.DialPassword("redis-01@20200417"),
	),
)

var sdk = New("****", "****", RedisStorage(rds))

func TestToken(t *testing.T) {
	tk1, err1 := sdk.token()
	tk2, err2 := sdk.token()
	assert.Nil(t, err1, err2)
	assert.NotEmpty(t, tk1, tk2)
	assert.EqualValues(t, tk1, tk2)
}
