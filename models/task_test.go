package main

import "testing"

func TestNewTask(t *testing.T) {
	title := "Make To-do app"
	task := NewTask(title)
	if task.Title != title {
		t.Errorf("expected title %q, got %q", title, task.Title)
	}
	if task.Done {
		t.Errorf("new task is done")
	}
}
