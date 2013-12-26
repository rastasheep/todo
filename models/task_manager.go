package main

type TaskManager struct {
	tasks []*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (m *TaskManager) Tasks() []*Task {
	return m.tasks
}

func (m *TaskManager) Add(task *Task) {
	m.tasks = append(m.tasks, task)
}
