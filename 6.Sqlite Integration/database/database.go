// database.go
package main.go

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the tasks table if it doesn't exist
	_, err = db.Exec(`
      CREATE TABLE IF NOT EXISTS tasks (
         id INTEGER PRIMARY KEY AUTOINCREMENT,
         title TEXT,
         description TEXT
      );
   `)
	if err != nil {
		log.Fatal(err)
	}
}

// Task structure
type Task struct {
	ID          int
	Title       string
	Description string
}

// Create a new task
func createTask(title, description string) error {
	_, err := db.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", title, description)
	return err
}

// Read all tasks
func getTasks() ([]Task, error) {
	rows, err := db.Query("SELECT id, title, description FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Update a task
func updateTask(id int, title, description string) error {
	_, err := db.Exec("UPDATE tasks SET title = ?, description = ? WHERE id = ?", title, description, id)
	return err
}

// Delete a task
func deleteTask(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
