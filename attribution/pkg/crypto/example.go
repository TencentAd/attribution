package main

import (
	"context"
	"fmt"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/conf"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/uitl"
	goredislib "github.com/go-redis/redis/v8"
	"strings"

	"math/big"
	"time"
)

func main() {
	encTest()
	//redisGetTest()
	//splitTest()
}

func encTest() {
	// 初始化加密相关的配置
	encryptConf := new(conf.EncryptConf)
	encryptConf.InitConf()

	redisConf, err := conf.NewRedisConf(conf.RedisCluster, "9.134.67.201:7000", "9.134.67.201:7005")
	//redisConf, err := conf.NewRedisConf(conf.RedisSingleNode, "9.134.67.201:6379")
	if err == conf.ErrNotSupportedRedisMode {
		fmt.Println("fail to create redis connection, please use supported redis mode")
		return
	}

	manager := uitl.NewKeyManager(redisConf)

	groupId := "groupId123"
	// 获取到加密使用的秘钥
	key, err := manager.GetEncryptKey(groupId)
	if err != nil {
		fmt.Println("fail to get the encrypt key")
		return
	}
	// 获取到加密使用的素数
	prime := new(big.Int)
	prime, _ = prime.SetString(encryptConf.Prime, 16)

	modPower := uitl.NewModPower(prime, key)
	encrypt := modPower.Encrypt(big.NewInt(123456))
	fmt.Println("encrypted data: ", encrypt.Text(62))
	decrypt := modPower.Decrypt(encrypt)
	fmt.Print("original data: ", decrypt)
}

func redisGetTest() {
	client := goredislib.NewClusterClient(&goredislib.ClusterOptions{
		Addrs: []string{
			"9.134.67.201:7000",
			"9.134.67.201:7001",
		},
		// TODO 将这个的配置加到配置文件中
		DialTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
	result, err := client.SetNX(context.Background(), "zefanliu1", "123", 0).Result()

	fmt.Println(result, err)
}

func splitTest() {
	s := string("111")
	split := strings.Split(s, ",")
	fmt.Println(split)
}
