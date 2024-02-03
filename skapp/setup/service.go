/*
 * @Descripttion: 初始化服务
 * @Author:DW
 * @Date: 2024-02-02 20:20:05
 * @LastEditTime: 2024-02-02 21:36:52
 */
package setup

import (
	"flag"
	"log"
	"secskill/skapp/endpoint"
	"secskill/skapp/service"
	"time"

	"golang.org/x/time/rate"
)

func InitServer(host string, servicePort string) {
	log.Printf("port is ", servicePort)

	flag.Parse()

	errChan := make(chan error)
	ratebucket := rate.NewLimiter(rate.Every(time.Second*1), 5000)
	var (
		skAppService service.Service
	)
	skAppService = service.SkAppService{}

	healthCheckEnd := endpoint.MakeHealthCheckEndpoint(skAppService)
	healthCheckEnd = 

}
