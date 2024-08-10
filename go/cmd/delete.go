package cmd

import (
	"strconv"

	"example.org/tuido/todos"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(delete)
}

var delete = &cobra.Command{
	Use:   "delete [number]",
	Short: "Delete a ToDo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		number, err := strconv.Atoi(args[0])

		if err == nil {
			todos.Delete(uint(number))
		}
	},
}
