package main

import "fmt"

type Task struct {
	Id    int64
	Title string
	Done  bool
}

func NewTask(title string) (*Task, error) {
	if len(title) == 0 {
		return nil, fmt.Errorf("empty title")
	}
	return &Task{0, title, false}, nil
}

func (task *Task) clone() *Task {
	clone := *task
	return &clone
}
