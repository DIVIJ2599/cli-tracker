package main

import (
	"fmt"
)

func main() {
	fmt.Println("Todo List")

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlag := NewCmdFlag()
	cmdFlag.Execute(&todos)
	storage.Save(todos)
}
