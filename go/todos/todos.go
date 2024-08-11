package todos

import (
	"fmt"
	"os"

	"example.org/tuido/todos/todo"
	"github.com/gocarina/gocsv"
	"github.com/jedib0t/go-pretty/v6/table"
)

type todos struct {
	path   string
	todos  []*todo.ToDo
	number uint
}

func new(path string) todos {
	tds, number := loadFromFile(path)

	return todos{
		path:   path,
		todos:  tds,
		number: number,
	}
}

func loadFromFile(path string) ([]*todo.ToDo, uint) {
	file, err := os.Open(path)
	defer file.Close()

	if err == nil {
		return fromCsv(file)
	} else {
		fmt.Printf("Error while reading File: %s\n", err)
		fmt.Printf("Creating new File:  %s\n", path)

		file, err := os.Create(path)
		defer file.Close()

		if err != nil {
			panic(fmt.Sprintf("Error while creating File: %s", err))
		}

		return make([]*todo.ToDo, 0), 1
	}

}

func fromCsv(file *os.File) ([]*todo.ToDo, uint) {
	var todos []*todo.ToDo

	if err := gocsv.UnmarshalFile(file, &todos); err != nil {
		fmt.Printf("Error while parsing ToDos: %s\n", err)
		return todos, 1
	}

	var number uint = 0
	for _, todo := range todos {
		if todo.Number > number {
			number = todo.Number
		}
	}

	number += 1

	return todos, number
}

func (tds *todos) write() {
	file, err := os.OpenFile(tds.path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	defer file.Close()

	if err != nil {
		panic(fmt.Sprintf("Error while opening File: %s", err))
	}

	err = gocsv.MarshalFile(&tds.todos, file)

	if err != nil {
		panic(fmt.Sprintf("Error while writing File: %s", err))
	}
}

func (tds *todos) add(content string) {
	todo := todo.New(tds.number, content)
	tds.todos = append(tds.todos, &todo)
}

func (tds *todos) complete(number uint) {
	for _, td := range tds.todos {
		if td.Number == number {
			td.Complete = true
			break
		}
	}
}

func (tds *todos) delete(number uint) {
	for i, td := range tds.todos {
		if td.Number == number {
			tds.todos = append(tds.todos[:i], tds.todos[i+1:]...)
			break
		}
	}
}

func (tds *todos) list(all bool, complete bool) {
	t := table.NewWriter()
	t.AppendHeader(todo.Titles())

	var filter func(todo *todo.ToDo) bool
	if all {
		filter = func(todo *todo.ToDo) bool { return true }
	} else if complete {
		filter = func(todo *todo.ToDo) bool { return todo.Complete }
	} else {
		filter = func(todo *todo.ToDo) bool { return !todo.Complete }
	}

	for _, td := range tds.todos {
		if filter(td) {
			t.AppendRow(td.Row())
		}
	}

	fmt.Print(t.Render())
}
