package cmd

import (
	"example.org/tuido/todos"
	"github.com/spf13/cobra"
)

var All bool
var Complete bool

func init() {
	rootCmd.AddCommand(list)
	list.Flags().BoolVarP(&All, "all", "a", false, "Display all ToDos")
	list.Flags().BoolVarP(&Complete, "complete", "c", false, "Display complete ToDos")
}

var list = &cobra.Command{
	Use:   "list",
	Short: "List all ToDos",
	Run: func(cmd *cobra.Command, args []string) {
		todos.List(All, Complete)
	},
}
