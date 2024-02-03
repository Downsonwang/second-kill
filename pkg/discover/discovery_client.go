/*
 * @Descripttion: 服务注册与发现客户端
 * @Author: DW
 * @Date: 2024-01-12 21:43:40
 * @LastEditTime: 2024-01-31 19:23:56
 */
package discover

import (
	"log"
	"secskill/pkg/common"
	"sync"

	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
)

type DiscoveryClientInstance struct {
	Host         string // HOST
	Port         int    // Port
	config       *api.Config
	client       consul.Client
	mutex        sync.Mutex
	instancesMap sync.Map
}
type DiscoveryClient interface {
	// 服务注册接口
	Register(instanceId, svcHost, healthCheckUrl, svcPort string, svcName string, weight int, meta map[string]string,
		tags []string, logger *log.Logger) bool

	// 服务注销接口
	DeRegister(instanceId string, logger *log.Logger) bool

	// 服务发现接口
	DiscoverServices(serviceName string, logger *log.Logger) []*common.ServiceInstance
}
