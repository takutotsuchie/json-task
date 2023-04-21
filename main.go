package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var name string
	var id int
	taskList := readFile()
	switch os.Args[1] {
	case "create":
		name = os.Args[2]
		// taskListにtaskを追加。
		id = len(taskList.List) + 1
		taskList.List = append(taskList.List, Task{id, name})
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if id <= 0 || id >= len(taskList.List)+1 {
			fmt.Println("無効なIDです")
			return
		}
		if err != nil {
			log.Print(err)
		}
		taskList.List = append(taskList.List[:id-1], taskList.List[id:]...)
		for i := 0; i < len(taskList.List); i++ {
			taskList.List[i].Id = i + 1
		}
		fmt.Println("Task List")
		for _, task := range taskList.List {
			fmt.Println(task.Id, task.Name)
		}
	case "ls":
		fmt.Println("Task List")
		for _, task := range taskList.List {
			fmt.Println(task.Id, task.Name)
		}
		return
	default:
		fmt.Println("command not found: ", os.Args[2])
		return
	}
	writeFile(taskList)
}
