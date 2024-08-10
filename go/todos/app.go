package todos

import (
	"fmt"

	"example.org/tuido/todos/todo"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Add(content string) {
	fmt.Printf("add %s\n", content)
}

func Complete(number uint) {
	fmt.Printf("complete %d\n", number)
}

func Delete(number uint) {
	fmt.Printf("delete %d\n", number)
}

func List(all bool, complete bool) {
	fmt.Print("list")

	if all {
		fmt.Print(" all")
	}

	if complete {
		fmt.Print(" complete")
	}

	fmt.Print("\n")

	t := table.NewWriter()
	t.AppendHeader(todo.Titles())

	td := todo.New(1, "My ToDo")
	t.AppendRow(td.Row())

	fmt.Print(t.Render())

}
