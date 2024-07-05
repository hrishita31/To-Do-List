package opr1

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func AddItem() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	ctx := context.Background()

	// var taskNo string
	// fmt.Println("Enter task number: ")
	// fmt.Scan(&taskNo)

	var task string
	fmt.Println("Enter task: ")
	fmt.Scan(&task)

	//(using string)
	// err := client.Set(ctx, taskNo, task, 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println("added task: ", task)

	//(using set)
	// err := client.SAdd(ctx, "task_set", task).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := client.Get(context.Background(), taskNo).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(taskNo, "-", val)

	var taskDone bool

	//(using hash)
	err := client.HSet(ctx, "task_hash", task, taskDone).Err()
	if err != nil {
		panic(err)
	}

	task, err = client.HGet(ctx, "task_hash", task).Result()
	if err != nil {
		panic(err)
	}
	//fmt.Println("Added task: ", task)

}
