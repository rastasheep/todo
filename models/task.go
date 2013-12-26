package main

import "fmt"

type Task struct {
	Title string
	Done  bool
}

func NewTask(title string) (*Task, error) {
	if len(title) == 0 {
		return nil, fmt.Errorf("empty title")
	}
	return &Task{title, false}, nil
}
