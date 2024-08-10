package cmd

import (
	"strconv"

	"example.org/tuido/todos"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(complete)
}

var complete = &cobra.Command{
	Use:   "complete [number]",
	Short: "Complete a ToDo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		number, err := strconv.Atoi(args[0])

		if err == nil {
			todos.Complete(uint(number))
		}
	},
}
