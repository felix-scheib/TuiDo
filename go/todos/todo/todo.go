package todo

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mergestat/timediff"
)

type ToDo struct {
	Number   uint      `csv:"number"`
	Content  string    `csv:"content"`
	Complete bool      `csv:"complete"`
	Date     time.Time `csv:"date"`
}

func New(number uint, content string) ToDo {
	return ToDo{
		Number:   number,
		Content:  content,
		Complete: false,
		Date:     time.Now(),
	}
}

func Titles() table.Row {
	return table.Row{
		"Number",
		"Content",
		"Complete",
		"Date",
	}
}

func (todo *ToDo) Row() table.Row {
	var complete string

	switch todo.Complete {
	case true:
		complete = "\u2714"
	case false:
		complete = "\u274C"
	}

	return table.Row{
		todo.Number,
		todo.Content,
		complete,
		timediff.TimeDiff(todo.Date),
	}
}
