package main

import (
	"fmt"
	"strconv"
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
			title := utils.GetUserInput("Provide title")
			description := utils.GetUserInput("Provide description")
			dueDate := utils.GetUserInput("Provide due date")
			taskManager.AddTask(&task.Task{ID: 1, Title: title, Description: description, DueDate: dueDate, Status: "NEW"})
		case "2":
			tasks := taskManager.ViewTasks()
			for _, task := range tasks {
				fmt.Printf("ID: %d, Title: %s\n", task.ID, task.Title)
			}

		case "3":
			title := utils.GetUserInput("Provide title")
			description := utils.GetUserInput("Provide description")
			dueDate := utils.GetUserInput("Provide due date")
			taskManager.UpdateTask(1, &task.Task{ID: 1, Title: title, Description: description, DueDate: dueDate, Status: "NEW"})
		case "4":
			inputId := utils.GetUserInput("Provide id")
			id, _ := strconv.Atoi(inputId)
			taskManager.DeleteTask(id)
		case "5":
			fmt.Println("Exiting Task Management System. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
