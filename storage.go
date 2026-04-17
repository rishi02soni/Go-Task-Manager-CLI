package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
	var tasks []Task

	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}

func getNextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func addTask(title string) {
	tasks, _ := loadTasks()

	task := Task{
		ID:    getNextID(tasks),
		Title: title,
		Done:  false,
	}

	tasks = append(tasks, task)
	saveTasks(tasks)

	fmt.Println("✅ Task added successfully!")
}

func listTasks() {
	tasks, _ := loadTasks()

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, t := range tasks {
		status := "❌"
		if t.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", t.ID, t.Title, status)
	}
}

func deleteTask(id int) {
	tasks, _ := loadTasks()
	var updated []Task

	for _, t := range tasks {
		if t.ID != id {
			updated = append(updated, t)
		}
	}

	saveTasks(updated)
	fmt.Println("🗑 Task deleted!")
}
