package main

import "testing"

func TestNewTaskManager(t *testing.T) {
	m := NewTaskManager()
	all := m.Tasks()

	if len(all) != 0 {
		t.Errorf("expected empty task manager")
	}
}

func TestAddAndRetrieve(t *testing.T) {
	task, err := NewTask("learn Go")
	if err != nil {
		t.Errorf("new task: %v", err)
	}

	m := NewTaskManager()
	m.Add(task)

	all := m.Tasks()
	if len(all) != 1 {
		t.Errorf("expected 1 task, got %v", len(all))
	}
	if *all[0] != *task {
		t.Errorf("expected %v, got %v", task, all[0])
	}
}
