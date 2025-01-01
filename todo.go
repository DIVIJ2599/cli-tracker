package main

import (
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

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index > len(*todos) {
		err := errors.New("index out of bounds")
		fmt.Println("Out of bounds: ", index)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//Delete the index
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		fmt.Println(err)
		return err
	}

	isCompleted := (*todos)[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		(*todos)[index].CompletedAt = &completionTime
	}

	(*todos)[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	err := todos.validateIndex(index)
	if err != nil {
		fmt.Println(err)
		return err
	}

	(*todos)[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for i, t := range *todos {
		completed := "X"
		completedAt := ""

		if t.Completed {
			completed = "V"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC3339)
			}
		}

		table.AddRow(strconv.Itoa(i), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
