/*
 * @Descripttion:
 * @Author:
 * @Date: 2024-01-12 21:39:16
 * @LastEditTime: 2024-02-01 19:54:48
 */
package common

// 服务依赖 + 基础结构
type ServiceInstance struct {
	Host      string //Host
	Port      int    // Port
	Weight    int    // 权重
	CurWeight int    // 当前权重

	GrpcPort int
}
