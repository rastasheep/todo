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

func TestAddAndRetrieveTwoTasks(t *testing.T) {
	learnGo, _ := NewTask("learn Go")
	learnTDD, _ := NewTask("learn TDD")

	m := NewTaskManager()
	m.Add(learnGo)
	m.Add(learnTDD)

	all := m.Tasks()
	if len(all) != 2 {
		t.Errorf("expected 2 tasks, got %v", len(all))
	}
	if *all[0] != *learnGo && *all[1] != *learnGo {
		t.Errorf("missing task: %v", learnGo)
	}
	if *all[0] != *learnTDD && *all[1] != *learnTDD {
		t.Errorf("missing task: %v", learnTDD)
	}
}

func TestAddModifyAndRetrieve(t *testing.T) {
	task, _ := NewTask("learn Go")
	m := NewTaskManager()
	m.Add(task)
	task.Done = true

	if m.Tasks()[0].Done {
		t.Errorf("saved task wasn't done")
	}
}

func TestAddAndFind(t *testing.T) {
	task, _ := NewTask("learn Go")
	m := NewTaskManager()
	m.Add(task)

	nt, ok := m.Find(task.Id)
	if !ok {
		t.Errorf("didn't find task")
	}
	if *task != *nt {
		t.Errorf("expected %v, got %v", task, nt)
	}
}
