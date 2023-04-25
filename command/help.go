package command

import "fmt"

func Help() {
	fmt.Println("help:")
	fmt.Println("     please check your command")
	fmt.Println("		create 'your task name'")
	fmt.Println("		delete 'your task ID'")
	fmt.Println("		ls")
}
