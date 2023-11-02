package task

type TaskManager struct {
	tasks []*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make([]*Task, 0),
	}
}

func (tm *TaskManager) AddTask(task *Task) {
	// Implement adding a task logic
}

func (tm *TaskManager) ViewTasks() []*Task {
	// Implement viewing tasks logic
	return tm.tasks
}

func (tm *TaskManager) UpdateTask(taskID int, updatedTask *Task) bool {
	// Implement updating a task logic
	return true
}

func (tm *TaskManager) DeleteTask(taskID int) bool {
	// Implement deleting a task logic
	return true
}
