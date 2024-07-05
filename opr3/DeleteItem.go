package opr3

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func DeleteItem() {

	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	ctx := context.Background()
	var itemToDel string
	fmt.Println("Enter the item you want to delete: ")
	fmt.Scan(&itemToDel)
	// err := client.Del(ctx, itemToDel).Err()
	// if err != nil {
	// 	panic(err)
	// }

	//(using set)
	// err := client.SRem(ctx, "task_set", itemToDel).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Deleted task: ", itemToDel)

	//using hash
	err := client.HDel(ctx, "task_hash", itemToDel).Err()
	if err != nil {
		panic(err)
	}
}
