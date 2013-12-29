package models

import "fmt"

type TaskManager struct {
	tasks  []*Task
	lastID int64
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (m *TaskManager) Tasks() []*Task {
	return m.tasks
}

func (m *TaskManager) Add(task *Task) error {
	if task.Id == 0 {
		m.lastID++
		task.Id = m.lastID
		m.tasks = append(m.tasks, task.clone())
		return nil
	}

	_, found := m.Find(task.Id)
	if found {
		m.tasks[task.Id-1] = task
		return nil
	}

	return fmt.Errorf("unknown task")
}

func (m *TaskManager) Find(id int64) (*Task, bool) {
	for _, t := range m.tasks {
		if t.Id == id {
			return t, true
		}
	}
	return nil, false
}
