package command

import (
	"fmt"

	"github.com/takutotsuchie/json-task/file"
	"github.com/takutotsuchie/json-task/types"
)

func Delete(id int, list []types.Task) {
	list = append(list[:id-1], list[id:]...)
	for i := 0; i < len(list); i++ {
		list[i].Id = i + 1
	}
	file.WriteFile(list)
	fmt.Println("Task List")
	for _, task := range list {
		fmt.Println(task.Id, task.Name)
	}
}
