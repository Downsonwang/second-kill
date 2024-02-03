/*
 * @Descripttion:
 * @Author:
 * @Date: 2024-02-02 20:48:43
 * @LastEditTime: 2024-02-02 20:48:47
 */
package bootstrap

var (
	HttpConfig         HttpConf
	DiscoverConfig     DiscoverConf
	ConfigServerConfig ConfigServerConf
	RpcConfig          RpcConf
)

// http配置
type HttpConf struct {
	Host string
	Port string
}

// RPC配置

type RpcConf struct {
	Port string
}

// 服务发现与注册配置
type DiscoverConf struct {
	Host        string
	Port        string
	ServiceName string
	Weight      int
	InstanceId  string
}

// 配置中心

type ConfigServerConf struct {
	Id      string
	Profile string
	Label   string
}
