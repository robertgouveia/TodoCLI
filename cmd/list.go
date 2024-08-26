package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var num = 0

var todoList = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		if num == 0 {
			fmt.Println("List all todos")
		} else {
			fmt.Println("List first", num, "todos")
		}
	},
}

func init() {
	todoList.Flags().IntVarP(&num, "num", "n", 0, "Number of Todos")

	rootCmd.AddCommand(todoList)
}
