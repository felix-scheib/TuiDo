package cmd

import (
	"fmt"
	"strconv"

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
		_, err := strconv.Atoi(args[0])

		if err != nil {
			return
		}

		fmt.Println("delete " + args[0])
	},
}
