package command

import (
	"github.com/takutotsuchie/json-task/file"
	"github.com/takutotsuchie/json-task/types"
)

func Create(name string, list []types.Task) {
	id := len(list) + 1
	list = append(list, types.Task{Id: id, Name: name})
	file.WriteFile(list)
}
