package cmd

import (
	"example.org/tuido/todos"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add [content]",
	Short: "Add a ToDo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todos.Add(args[0])
	},
}
