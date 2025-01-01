package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type cmdFlag struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlag() *cmdFlag {
	cf := cmdFlag{}

	flag.StringVar(&cf.Add, "add", "", "Add a task")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task by index. Example: id:<NEW_TITLE>")
	flag.IntVar(&cf.Del, "delete", -1, "Specify a task to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a task")
	flag.BoolVar(&cf.List, "list", false, "List all tasks")

	flag.Parse()
	return &cf
}

func (cf *cmdFlag) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.List:
		todos.print()
	case cf.Toggle != -1:
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit command. Example: id:<NEW_TITLE>")
			return
		}
		id, err := strconv.Atoi(parts[0][3:])
		if err != nil || id < 0 || id > len(*todos)-1 {
			fmt.Println("Invalid task index")
			return
		}
		todos.edit(id, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	}
}
