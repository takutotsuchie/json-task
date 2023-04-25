package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/takutotsuchie/json-task/command"
	"github.com/takutotsuchie/json-task/file"
)

func main() {

	taskList := file.ReadFile()
	if len(os.Args) == 1 {
		command.Help()
		return
	}
	switch os.Args[1] {
	case "create":
		name := os.Args[2]
		// taskListにtaskを追加。
		command.Create(name, taskList)
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Print(err)
		}
		if id <= 0 || id >= len(taskList)+1 {
			fmt.Println("無効なIDです")
			return
		}
		command.Delete(id, taskList)

	case "ls":
		fmt.Println("Task List")
		for _, task := range taskList {
			fmt.Println(task.Id, task.Name)
		}
		return
	default:
		command.Help()
		return
	}

}
