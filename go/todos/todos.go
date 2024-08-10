package todos

import (
	"fmt"
	"os"

	"example.org/tuido/todos/todo"
	"github.com/gocarina/gocsv"
)

type todos struct {
	path   string
	todos  []todo.ToDo
	number uint
}

func new(path string) todos {
	loadFromFile(path)

	return todos{
		path:   path,
		todos:  []todo.ToDo{},
		number: 0,
	}
}

func loadFromFile(path string) ([]todo.ToDo, uint) {
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

		return make([]todo.ToDo, 0), 1
	}

}

func fromCsv(file *os.File) ([]todo.ToDo, uint) {
	return make([]todo.ToDo, 0), 1
}

func (tds *todos) Write() {
	todos := make([]todo.ToDo, 0)
	todos = append(todos, todo.New(1, "First ToDo"))
	todos = append(todos, todo.New(2, "Second ToDo"))
	todos = append(todos, todo.New(3, "Third ToDo"))

	file, err := os.OpenFile("./tds.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()

	if err != nil {
		panic(fmt.Sprintf("Error while opening File: %s", err))
	}

	err = gocsv.MarshalFile(&todos, file)

	if err != nil {
		panic(fmt.Sprintf("Error while writing File: %s", err))
	}

	fmt.Println(len(todos))
}
