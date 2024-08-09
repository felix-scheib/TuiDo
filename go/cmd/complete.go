package cmd

import (
	"fmt"
	"strconv"

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
		_, err := strconv.Atoi(args[0])

		if err != nil {
			return
		}
		fmt.Println("complete " + args[0])
	},
}
