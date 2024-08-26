package cmd

import (
	"fmt"
	"strconv"

	"github.com/robertgouveia/TodoCLI/local"
	"github.com/spf13/cobra"
)

var updateBool bool

var updateTodo = &cobra.Command{
	Use:   "update",
	Short: "Update a todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		updateIndex, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid index")
			return
		}
		if updateIndex == 0 {
			fmt.Println("Please provide an index starting from 1")
			return
		}
		err = local.UpdateTodos(updateIndex, updateBool)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Todo updated")
	},
}

func init() {
	updateTodo.Flags().BoolVarP(&updateBool, "done", "d", false, "Set todo as done")
	rootCmd.AddCommand(updateTodo)
}
