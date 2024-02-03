/*
 * @Descripttion: LoadBalance
 * @Author:DW
 * @Date: 2024-01-31 23:51:46
 * @LastEditTime: 2024-02-01 00:31:06
 */
package loadbalance

import (
	"errors"
	"math/rand"
	"secskill/pkg/common"
)

// 完全随机策略 从服务实例列表中选择一个服务实例返回
type LoadBalance interface {
}

type RandomLoadBalance struct {
}

func (loadBalance *RandomLoadBalance) SelectService(service []*common.ServiceInstance) (*common.ServiceInstance, error) {
	if service == nil || len(service) == 0 {
		return nil, errors.New("service instances are not exist")
	}

	return service[rand.Intn(len(service))], nil
}

// 权重平滑负载均衡
type WeightRoundRobinLoadBalance struct {
}

func (loadBalance *WeightRoundRobinLoadBalance) SelectService(services []*common.ServiceInstance) (best *common.ServiceInstance, err error) {
	if services == nil || len(services) == 0 {
		return nil, errors.New("service instances are not exist ")
	}

	total := 0
	// 遍历所有服务实例
	for i := 0; i < len(services); i++ {
		w := services[i]
		if w == nil {
			continue
		}
		// CurWeight加权
		w.CurWeight += w.Weight
		// 
		total += w.Weight

		if best == nil || w.CurWeight > best.CurWeight {
			best = w
		}
	}

	if best == nil {
		return nil, nil
	}

	best.CurWeight -= total
	return best, nil
}
