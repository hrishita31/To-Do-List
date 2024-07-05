package main

import (
	"fmt"
	"os"
	"to-do/opr1"
	"to-do/opr2"
	"to-do/opr3"
	"to-do/opr4"
)

func main() {

	for true {
		fmt.Println("Enter 1 to add task")
		fmt.Println("Enter 2 to view all tasks")
		fmt.Println("Enter 3 to delete task")
		fmt.Println("Enter 4 to mark a task as done")
		fmt.Println("Enter 5 to exit")

		var n int
		fmt.Println("Enter choice: ")
		fmt.Scan(&n)
		switch n {
		case 1:
			opr1.AddItem()

		case 2:
			//view all items
			opr2.ViewAll()
		case 3:
			//deleteItem()
			opr3.DeleteItem()
		case 4:
			opr4.MarkAsDone()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Enter correct value")
			break
		}
	}
}
