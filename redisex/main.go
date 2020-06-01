package main

import (
	"fmt"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

func main() {

	// Connect to redis.
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "192.168.163.133:6379",
	})
	defer client.Close()
	client.Set("money",10000000000,0)
	for {
		time.Sleep(time.Second)
		// Create a new lock client.
		locker := redislock.New(client)

		// Try to obtain lock.
		lock, err := locker.Obtain("my-key", 100*time.Millisecond, nil)
		if err == redislock.ErrNotObtained {
			fmt.Println("Could not obtain lock!")
		} else if err != nil {
			log.Fatalln(err)
		}

		// Don't forget to defer Release.
		defer lock.Release()
		fmt.Println("I have a lock!")
		m,_ := client.Get("money").Int()
		fmt.Println("money:", m, " after:", m/2+6)
		client.Set("money", m/2+6, 0)
		// Sleep and check the remaining TTL.
		time.Sleep(50 * time.Millisecond)
		if ttl, err := lock.TTL(); err != nil {
			log.Fatalln(err)
		} else if ttl > 0 {
			fmt.Println("Yay, I still have my lock!")
		}

		// Extend my lock.
		if err := lock.Refresh(100*time.Millisecond, nil); err != nil {
			log.Fatalln(err)
		}

		// Sleep a little longer, then check.
		time.Sleep(100 * time.Millisecond)
		if ttl, err := lock.TTL(); err != nil {
			log.Fatalln(err)
		} else if ttl == 0 {
			fmt.Println("Now, my lock has expired!")
		}
	}
}
