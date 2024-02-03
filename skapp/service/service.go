/*
 * @Descripttion:define the service interface
 * @Author:DW
 * @Date: 2024-02-02 20:26:37
 * @LastEditTime: 2024-02-02 21:58:33
 */
package service

import (
	"secskill/skadmin/config"
	"secskill/skapp/model"
)

type Service interface {
	HealthCheck() bool
	SecInfo(productId int) (date map[string]interface{})
	SecKill(req *model.SecRequest) (map[string]interface{}, int, error)
	SecInfoList() ([]map[string]interface{}, int, error)
}

type SkAppService struct{}

// 用于检查服务的健康状态 这里仅仅返回true
func (s SkAppService) HealthCheck() bool {
	return true
}


type ServiceMiddleware func(Service) Service


func (s SkAppService) SecInfo(productId int) (date map[string]interface{}) {
	config.SkAppContext.RWSecProductLock.RLock()
	defer config.SkAppContext.RWSecProductLock.Unlock()
	v, ok := 
}

func (s SkAppService) SecInfoList() ([]map[string]interface{}, int, error) {

}