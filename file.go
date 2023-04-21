package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var file = os.Getenv("taskfile")

func readFile() TaskList {
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
	var list TaskList
	err = json.Unmarshal(data, &list)
	if err != nil {
		log.Print(err)
	}
	return list
}

func writeFile(list TaskList) {
	// ファイルを空にする。
	f, err := os.Create(file)
	if err != nil {
		log.Print(err)
	}
	// jsonにする
	data, err := json.Marshal(list)
	if err != nil {
		log.Print(err)
	}
	//書き込む。
	_, err = f.Write(data)
	if err != nil {
		log.Print(err)
	}
}
