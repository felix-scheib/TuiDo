package todo

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mergestat/timediff"
)

type ToDo struct {
	number   uint
	content  string
	complete bool
	date     time.Time
}

func New(number uint, content string) ToDo {
	return ToDo{
		number:   number,
		content:  content,
		complete: false,
		date:     time.Now(),
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
	if todo.complete {
		complete = "\u2714"
	} else {
		complete = "\u274C"
	}

	return table.Row{
		todo.number,
		todo.content,
		complete,
		timediff.TimeDiff(todo.date),
	}
}
