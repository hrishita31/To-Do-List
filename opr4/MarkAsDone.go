package opr4

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func MarkAsDone() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	ctx := context.Background()

	var task string
	fmt.Println("Enter the string to be marked as done: ")
	fmt.Scan(&task)

	// taskDone, err := client.HGet(ctx, "task_hash", task).Result()
	// if err != nil {
	// 	panic(err)
	// }

	//taskDone := fmt.Sprintf("%s-done", task)
	//taskDone = "1"

	//fmt.Println("Done task is: ", task, "-", taskDone)

	err := client.HSet(ctx, "task_hash", task, "1").Err()
	if err != nil {
		panic(err)
	}

}
