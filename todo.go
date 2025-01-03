package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aquasecurity/table"
)

type Status int

const (
	ToDo Status = iota
	InProgress
	Completed
)

func (s Status) String() string {
	switch s {
	case ToDo:
		return "Todo"
	case InProgress:
		return "In Progress"
	case Completed:
		return "Completed"
	default:
		return "Unknown"
	}
}

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	Deadline    time.Time
	OverDue     bool
	Status      Status
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
		OverDue:     false,
		Status:      ToDo,
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
		(*todos)[index].Status = Completed
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

func (todos *Todos) setDeadline(index int, deadline string) error {
	// Validate the index
	err := todos.validateIndex(index)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Parse the deadline
	var parsedDeadline time.Time
	if strings.HasSuffix(deadline, "d") { // e.g., "3d" for 3 days
		days, err := strconv.Atoi(strings.TrimSuffix(deadline, "d"))
		if err != nil || days < 0 {
			return fmt.Errorf("invalid deadline format: %s", deadline)
		}
		parsedDeadline = time.Now().Add(time.Hour * 24 * time.Duration(days))
	} else if strings.HasSuffix(deadline, "h") { // e.g., "5h" for 5 hours
		hours, err := strconv.Atoi(strings.TrimSuffix(deadline, "h"))
		if err != nil || hours < 0 {
			return fmt.Errorf("invalid deadline format: %s", deadline)
		}
		parsedDeadline = time.Now().Add(time.Hour * time.Duration(hours))
	} else { // Assume a specific date format: "YYYY-MM-DD" or "YYYY-MM-DD HH:mm"
		parsedDeadline, err = time.Parse("2006-01-02 15:04", deadline)
		if err != nil {
			parsedDeadline, err = time.Parse("2006-01-02", deadline)
			if err != nil {
				return fmt.Errorf("invalid deadline format: %s", deadline)
			}
		}
	}

	// Set the parsed deadline
	(*todos)[index].Deadline = parsedDeadline
	(*todos)[index].Status = InProgress
	fmt.Printf("Deadline for task %d set to %s\n", index, parsedDeadline.Format(time.RFC1123))
	return nil
}

func (todos *Todos) markInProgress(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if (*todos)[index].Status == InProgress {
		fmt.Printf("Task %d is already in progress\n", index)
		return nil
	}

	(*todos)[index].Status = InProgress

	return nil
}

func (todos *Todos) markCompleted(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if (*todos)[index].Status == Completed {
		fmt.Printf("Task %d is already Completed\n", index)
		return nil
	}

	(*todos)[index].Status = InProgress
	(*todos)[index].Completed = true
	completionTime := time.Now()
	(*todos)[index].CompletedAt = &completionTime

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Status", "Completed", "Created At", "Completed At", "Deadline", "Overdue")
	for i, t := range *todos {
		completed := "X"
		completedAt := ""
		overdue := "X"

		if t.Completed {
			completed = "V"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC3339)
			}
		}

		if t.Deadline.Before(time.Now()) {
			overdue = "V"
		}

		table.AddRow(strconv.Itoa(i), t.Title, t.Status.String(), completed, t.CreatedAt.Format(time.RFC1123), completedAt, t.Deadline.Format(time.RFC1123), overdue)
	}

	table.Render()
}
