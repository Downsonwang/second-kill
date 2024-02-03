package setup

import (
	"fmt"
	"log"
	"secskill/pkg/config"
	"time"

	"github.com/unknwon/com"

	"github.com/go-redis/redis"
)

// 初始化Redis
func InitRedis() {
	log.Printf("init redis %s", config.Redis.Password)
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("Connect redis failed. Error : %v", err)
	}

	log.Printf("init redis success")
	config.Redis.RedisConn = client

}

func loadBlackList(conn *redis.Client) {
	config.SecKill.IPBlackMap = make(map[string]bool, 10000)
	config.SecKill.IDBlackMap = make(map[int]bool, 10000)

	// 用户ID
	idList, err := conn.HGetAll(config.Redis.IdBlackListHash).Result()
	if err != nil {
		log.Printf("hget all failed. Error : %v", err)
		return
	}

	for _, v := range idList {
		id, err := com.StrTo(v).Int()
		if err != nil {
			log.Printf("invalid user id [%v]", id)
			continue
		}
		config.SecKill.IDBlackMap[id] = true
	}

	//用户IP
	ipList, err := conn.HGetAll(config.Redis.IdBlackListHash).Result()

	if err != nil {
		log.Printf("hget all failed. Error : %v", err)
		return
	}
	for _, v := range ipList {
		config.SecKill.IPBlackMap[v] = true
	}

	return
}

// 同步用户ID黑名单
func syncIdBlackList(conn *redis.Client) {
	for {
		// Redis 的 BRPOP 命令从队列中获取 ID，并将其添加到一个 ID 黑名单映射中
		idArr, err := conn.BRPop(time.Minute, config.Redis.IdBlackListQueue).Result()
		if err != nil {
			log.Printf("brpop id failed, err : %v", err)
			continue
		}

		id, _ := com.StrTo(idArr[1]).Int()
		// 写锁保护
		config.SecKill.RWBlackLock.Lock()
		{
			config.SecKill.IDBlackMap[id] = true
		}
		// 释放写锁
		config.SecKill.RWBlackLock.Unlock()
	}
}

// 同步用户IP黑名单
func syncIpBlackList(conn *redis.Client) {
	var ipList []string
	lastTime := time.Now().Unix()

	for {
		ipArr, err := conn.BRPop(time.Minute, config.Redis.IpBlackListQueue).Result()
		if err != nil {
			log.Printf("brpop ip failed ,err : %v", err)
			continue
		}
		ip := ipArr[1]
		curTime := time.Now().Unix()
		ipList = append(ipList, ip)

		if len(ipList) > 100 || curTime-lastTime > 5 {
			config.SecKill.RWBlackLock.Lock()
			{
				for _, v := range ipList {
					// 并将其值设为 true，表示该 IP 在黑名单中。
					config.SecKill.IPBlackMap[v] = true
				}
			}
			config.SecKill.RWBlackLock.Unlock()
			lastTime = curTime //记录当前IP添加的时间
			log.Printf("sync ip list from redis success, ip[%v]", ipList)

		}
	}
}

// 初始化redis进程

func initRedisProcess() {
	log.Printf("initRedisProcess %d %d", config.SecKill.AppWriteToHandleGoroutineNum, config.SecKill.AppReadFromHandleGoroutineNum)
	for i := 0; i < config.SecKill.AppWriteToHandleGoroutineNum; i++ {
		go srv_redis.WriteHandle()
	}

	for i := 0; i < config.SecKill.AppReadFromHandleGoroutineNum; i++ {
		go srv_redis.ReadHandle()
	}
}


