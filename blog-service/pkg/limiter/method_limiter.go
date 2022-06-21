package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

//实现接口限流器
type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterIface {
	return MethodLimiter{
		Limiter: &Limiter{
			limiterBuckets: make(map[string]*ratelimit.Bucket),
		},
	}
}

// 得到请求的uri
func (m MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (m MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := m.limiterBuckets[key]
	return bucket, ok
}

func (m MethodLimiter) AddBuckets(rules ...LimiterfaceBucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := m.limiterBuckets[rule.Key]; !ok {
			m.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}
	return m
}
