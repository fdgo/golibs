package main

import (
	"github.com/go-redis/redis"
	"log"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "120.27.239.127:6379",              // 使用默认数据库
		Password: "962012d09b8170d912f0669f6d7d9d07", // 没有密码则置空
		DB:       0,                                  // 使用默认的数据库
	})

	pong, err := rdb.Ping().Result() // 检查是否连接
	if err != nil {
		log.Fatal(err)
	}

	// 连接成功啦
	log.Println(pong)

	// 订阅全部消息
	pubsub := rdb.Subscribe("__keyevent@0__:expired")
	// 等待消息返回，原因是上一个方法不是立即返回的，囧
	_, err = pubsub.Receive()
	if err != nil {
		log.Fatal(err)
	}

	// 用管道来接收消息
	ch := pubsub.Channel()

	// 处理消息
	for msg := range ch {
		log.Println(msg.Channel, ":", msg.Payload) //__keyevent@0__:expired : mykey1   // __keyevent@0__:expired : mykey2
	}

}
// set mykey1 3  myval1
// set mykey2 3  myval2