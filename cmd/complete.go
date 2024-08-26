package cmd

import (
	"fmt"
	"strconv"

	"github.com/robertgouveia/TodoCLI/local"
	"github.com/spf13/cobra"
)

var taskIndex int

var completedTodo = &cobra.Command{
	Use:   "complete",
	Short: "Complete Todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskIndex, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid index")
			return
		}

		err = local.CompleteTodo(taskIndex)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Todo completed")
	},
}

func init() {
	rootCmd.AddCommand(completedTodo)
}
