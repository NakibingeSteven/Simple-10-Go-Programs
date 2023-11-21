package main.go

import (
	"fmt"
	"log"
)

func main() {
	// Create a task
	err := createTask("Task 1", "Description for Task 1")
	if err != nil {
		log.Fatal(err)
	}

	// Read all tasks
	tasks, err := getTasks()
	if err != nil {
		log.Fatal(err)
	}

	// Display tasks
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s\n", task.ID, task.Title, task.Description)
	}

	// Update a task
	err = updateTask(1, "Updated Task 1", "Updated Description for Task 1")
	if err != nil {
		log.Fatal(err)
	}

	// Read all tasks after update
	tasks, err = getTasks()
	if err != nil {
		log.Fatal(err)
	}

	// Display updated tasks
	fmt.Println("\nAfter Update:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s\n", task.ID, task.Title, task.Description)
	}

	// Delete a task
	err = deleteTask(1)
	if err != nil {
		log.Fatal(err)
	}

	// Read all tasks after delete
	tasks, err = getTasks()
	if err != nil {
		log.Fatal(err)
	}

	// Display remaining tasks
	fmt.Println("\nAfter Delete:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s\n", task.ID, task.Title, task.Description)
	}
}
