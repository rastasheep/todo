package main

import "testing"

func TestNewTask(t *testing.T) {
	title := "Make To-do app"
	task, err := NewTask(title)
	if task.Title != title {
		t.Errorf("expected title %q, got %q", title, task.Title)
	}
	if task.Done {
		t.Errorf("new task is done")
	}
	if err != nil {
		t.Errorf("task should be crated")
	}
}
func TestNewTaskEmptyTitle(t *testing.T) {
	_, err := NewTask("")
	if err == nil {
		t.Errorf("task with empty title created")
	}
}
