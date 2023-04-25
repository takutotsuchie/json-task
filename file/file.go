package file

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/takutotsuchie/json-task/types"
)

var file = os.Getenv("taskfile")

func ReadFile() []types.Task {
	//ファイルを開く
	f, err := os.Open(file)
	if err != nil {
		log.Print(err)
	}
	// ファイルを読む
	data, err := io.ReadAll(f)
	if err != nil {
		log.Print(err)
	}
	// 構造体にする
	var list []types.Task
	err = json.Unmarshal(data, &list)
	if err != nil {
		log.Print(err)
	}
	return list
}

func WriteFile(list []types.Task) {
	// ファイルを空にする。
	f, err := os.Create(file)
	if err != nil {
		log.Print(err)
	}
	// jsonにする
	data := jsonMarshaler(list)
	if err != nil {
		log.Print(err)
	}
	//書き込む。
	_, err = f.Write(data)
	if err != nil {
		log.Print(err)
	}
}

func jsonMarshaler(list []types.Task) []byte {
	var bytes []byte
	bytes = append(bytes, []byte("[\n")...)
	for index, t := range list {
		data, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}
		bytes = append(bytes, data...)
		if index != len(list)-1 {
			bytes = append(bytes, []byte(",\n")...)
		} else {
			bytes = append(bytes, []byte("\n")...)
		}
	}
	bytes = append(bytes, []byte("]\n")...)
	return bytes
}
