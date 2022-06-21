package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterfaceBucketRule) LimiterIface
}

// 令牌桶与键值对的映射关系
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// 存储令牌桶的规则属性
type LimiterfaceBucketRule struct {
	Key          string
	FillInterval time.Duration //填充间隔时间
	Capacity     int64
	Quantum      int64
}
