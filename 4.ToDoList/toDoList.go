package main

import (
	"fmt"
	"os"
)

const fileName = "tasks.txt"

func main() {
	// Load tasks from file
	tasks := loadTasks()

	for {
		fmt.Println("\nTo-Do List:")
		listTasks(tasks)

		fmt.Print("\nOptions:\n1. Add Task\n2. Remove Task\n3. Quit\nEnter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask(tasks)
		case 2:
			removeTask(tasks)
		case 3:
			saveTasks(tasks)
			fmt.Println("Quitting the To-Do List. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}

func listTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks.")
		return
	}

	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func addTask(tasks []string) {
	fmt.Print("Enter the task: ")
	var task string
	fmt.Scanln(&task)

	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Println("Task added successfully.")
}

func removeTask(tasks []string) {
	fmt.Print("Enter the task number to remove: ")
	var taskNum int
	fmt.Scanln(&taskNum)

	if taskNum < 1 || taskNum > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	taskToRemove := tasks[taskNum-1]
	tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)
	saveTasks(tasks)
	fmt.Printf("Task '%s' removed successfully.\n", taskToRemove)
}

func loadTasks() []string {
	file, err := os.Open(fileName)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var tasks []string
	for {
		var task string
		_, err := fmt.Fscanln(file, &task)
		if err != nil {
			break
		}
		tasks = append(tasks, task)
	}

	return tasks
}

func saveTasks(tasks []string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		fmt.Fprintln(file, task)
	}
}
