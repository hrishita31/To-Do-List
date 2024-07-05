package opr2

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ViewAll() {

	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	ctx := context.Background()
	//(using set)
	// tasks, err := client.SMembers(ctx, "task_set").Result()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//using hash
	tasks, err := client.HGetAll(ctx, "task_hash").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("All tasks are : ", tasks)
}
