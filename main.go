package main

import (
	"fmt"
	"task-manager/task"
	"task-manager/utils"
)

func main() {
	taskManager := task.NewTaskManager()

	for {
		fmt.Println("Task Management System")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Update Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		choice := utils.GetUserInput("Enter your choice")

		switch choice {
		case "1":
			taskManager.AddTask(&task.Task{ID: 1, Title: "", Description: "", DueDate: "", Status: ""})
		case "2":
			// Logic to view tasks
		case "3":
			// Logic to update a task
		case "4":
			// Logic to delete a task
		case "5":
			fmt.Println("Exiting Task Management System. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
