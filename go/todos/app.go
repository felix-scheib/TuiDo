package todos

const PATH string = "./todos.csv"

func Add(content string) {
	todos := new(PATH)
	defer todos.write()

	todos.add(content)
}

func Complete(number uint) {
	todos := new(PATH)
	defer todos.write()

	todos.complete(number)
}

func Delete(number uint) {
	todos := new(PATH)
	defer todos.write()

	todos.delete(number)
}

func List(all bool, complete bool) {
	todos := new(PATH)

	todos.list(all, complete)
}
