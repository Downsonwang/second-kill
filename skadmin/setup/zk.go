package setup

import (
	"encoding/json"
	"fmt"
	"log"
	"secskill/pkg/config"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func InitZk() {
	var hosts = []string{"127.0.0.1:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	config.Zk.ZkConn = conn
	config.Zk.SecProductKey = "/product"

}

// 加载秒杀信息
func loadSecConf(conn *zk.Conn) {
	log.Println("connect zk success %s", config.Zk.SecProductKey)

	v, _, err := conn.Get(config.Zk.SecProductKey)
	if err != nil {
		log.Printf("get product info failed,err:%v", err)
		return
	}
	log.Printf("get product info ")
	var secProductInfo []*config.SecProductInfoConf
	err = json.Unmarshal(v, &secProductInfo)
	if err != nil {
		log.Printf("Unmsharl second product info failed, err : %v", err)
	}
	updateSecPorductInfo(secProductInfo)
}

// 当有数值发生变化时,会调用waitSecProductEvent方法 ,传入对应的zk.Event
// Event中会携带本次事件涉及的路径和类型和状态
func waitSecProductEvent(event zk.Event) {
	log.Print(">>>>>>>>>>>>>>>>>>>")
	log.Println("path:", event.Path)
	log.Println("type:", event.Type.String())
	log.Println("state:", event.State.String())
	log.Println("<<<<<<<<<<<<<<<<<<<")
	if event.Path == config.Zk.SecProductKey {

	}
}

// 更新秒杀商品信息

func updateSecPorductInfo(secProductInfo []*config.SecProductInfoConf) {
	tmp := make(map[int]*config.SecProductInfoConf, 1024)
	for _, v := range secProductInfo {
		log.Printf("updateSecProductInfo %v", v)
		tmp[v.ProductId] = v
	}
	config.SecKill.RWBlackLock.Lock()
	config.SecKill.SecProductInfoMap = tmp
	config.SecKill.RWBlackLock.Unlock()

}
