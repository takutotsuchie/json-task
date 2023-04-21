package main

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type TaskList struct {
	List []Task `json:"tasks"`
}
