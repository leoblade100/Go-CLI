package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// make []Todo everywhere ---> Todos (easier to read)
type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) togle(index int) error {
	t := *todos
	Completedtion_time := time.Now()

	if err := t.validateIndex(index); err != nil {
		return err
	}

	if t[index].Completed != true {
		t[index].Completed = true
		t[index].CompletedAt = &Completedtion_time
	} else {
		t[index].Completed = false
		t[index].CompletedAt = nil

	}
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	t[index].Title = title
	return nil
}

func (todos *Todos) saveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Serialize Todos into JSON
	encoder := json.NewEncoder(file)
	return encoder.Encode(todos)
}

func (todos *Todos) loadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, return nil (nothing to load)
			return nil
		}
		return err
	}
	defer file.Close()

	// Deserialize JSON into Todos
	decoder := json.NewDecoder(file)
	return decoder.Decode(todos)
}

func (todos *Todos) print() error {
	a := *todos

	t := table.New(os.Stdout)
	t.SetRowLines(false)
	t.SetHeaders("#", "Title", "Completed", "Create at", "Completed at")

	for i, todo := range a {
		createdAt := todo.CreatedAt.Format(time.RFC822)
		completedAt := "Not completed"
		emoji := "❌"

		// Handle nil CompletedAt field
		if todo.CompletedAt != nil {
			emoji = "✅"
			completedAt = todo.CompletedAt.Format(time.RFC822)
		}

		// Add row with formatted values
		t.AddRow(
			strconv.Itoa(i), // Index
			todo.Title,      // Title
			emoji,           // Completed
			createdAt,       // Created at
			completedAt,     // Completed at
		)
	}
	t.Render()
	return nil
}
