package task

import (
	"encoding/json"
	"os"
)

const filePath = "storage.json"

type TaskManager struct {
	tasks []*Task
}

func NewTaskManager() *TaskManager {
	content, err := os.ReadFile(filePath)
	PanicIfError(err)
	var tasks []*Task
	json.Unmarshal(content, &tasks)
	return &TaskManager{
		tasks: tasks,
	}
}

func (tm *TaskManager) AddTask(task *Task) {
	tasks := append(tm.tasks, task)
	data, _ := json.Marshal(tasks)
	err := os.WriteFile(filePath, data, 0644)
	PanicIfError(err)
}

func (tm *TaskManager) ViewTasks() []*Task {
	return tm.tasks
}

func (tm *TaskManager) UpdateTask(taskID int, updatedTask *Task) bool {
	task, exists := findTaskByID(tm.tasks, taskID)
	if !exists {
		return false
	}
	task.Description = updatedTask.Description
	task.Title = updatedTask.Title
	task.Status = updatedTask.Status
	task.DueDate = updatedTask.DueDate
	return true
}

func findTaskByID(tasks []*Task, targetID int) (*Task, bool) {
	for _, task := range tasks {
		if task.ID == targetID {
			return task, true
		}
	}

	return nil, false
}

func (tm *TaskManager) DeleteTask(taskID int) bool {
	var updatedTasks []*Task

	for _, task := range tm.tasks {
		if task.ID != taskID {
			updatedTasks = append(updatedTasks, task)
		}
	}
	if len(tm.tasks) == len(updatedTasks) {
		return false
	}
	tm.tasks = updatedTasks
	data, _ := json.Marshal(updatedTasks)
	os.WriteFile(filePath, data, 0644)
	return true
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
