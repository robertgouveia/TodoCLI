package cmd

import (
	"fmt"
	"strconv"

	"github.com/robertgouveia/TodoCLI/local"
	"github.com/spf13/cobra"
)

var index int

var removeTodo = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid index")
			return
		}

		err = local.RemoveTodos(index)
		if err != nil {
			fmt.Println("Error removing todo:", err)
			return
		}

		fmt.Println("Todo removed successfully")
	},
}

func init() {
	rootCmd.AddCommand(removeTodo)
}
