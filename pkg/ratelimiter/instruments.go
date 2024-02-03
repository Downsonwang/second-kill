/*
 * @Descripttion:限流器
 * @Author:DW
 * @Date: 2024-02-01 10:54:27
 * @LastEditTime: 2024-02-01 11:05:05
 */

package ratelimiter

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

var ErrLimitExceed = errors.New("Rate limit exceed!")

// 创建限流中间件
func NewTokenBucketLimitterWithBuildIn(bkt *rate.Limiter) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !bkt.Allow() {
				return nil, ErrLimitExceed
			}
			return e(ctx, request)
		}
	}
}

func DynamicLimitter(interval int, burst int) endpoint.Middleware {
	// 创建一个令牌桶，限流时间间隔为 interval 秒，允许的突发请求数量为 burst
	bucket := rate.NewLimiter(rate.Every(time.Second*time.Duration(interval)), burst)

	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !bucket.Allow() {
				return nil, ErrLimitExceed
			}
			return e(ctx, request)
		}
	}

}
