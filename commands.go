package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlag struct {
	Add            string
	Del            int
	Edit           string
	Toggle         int
	List           bool
	SetDeadline    string
	MarkInProgress int
	MarkCompleted  int
}

func NewCmdFlag() *cmdFlag {
	cf := cmdFlag{}

	flag.StringVar(&cf.Add, "add", "", "Add a task")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task by index. Example: id:<NEW_TITLE>")
	flag.IntVar(&cf.Del, "delete", -1, "Specify a task to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a task")
	flag.BoolVar(&cf.List, "list", false, "List all tasks")
	flag.StringVar(&cf.SetDeadline, "setDeadline", "", "Specify the deadline")
	flag.IntVar(&cf.MarkInProgress, "markInProgress", -1, "Mark a task as in progress")
	flag.IntVar(&cf.MarkCompleted, "markCompleted", -1, "Mark a task as completed")
	flag.Parse()
	return &cf
}

func (cf *cmdFlag) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.List:
		todos.print()
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		fmt.Println(len(parts))
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.SetDeadline != "":
		parts := strings.SplitN(cf.SetDeadline, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: invalid format for Setting Deadline. Use id:deadline (e.g., 1:3d or 1:2025-01-01)")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid task index")
			os.Exit(1)
		}

		err = todos.setDeadline(index, parts[1])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case cf.MarkInProgress != -1:
		todos.markInProgress(cf.MarkInProgress)
	case cf.MarkCompleted != -1:
		todos.markCompleted(cf.MarkCompleted)
	}
}
