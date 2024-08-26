package cmd

import (
	"fmt"
	"time"

	"github.com/robertgouveia/TodoCLI/local"
	"github.com/spf13/cobra"
)

var title string

var add = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var todo local.Todo

		title = args[0]

		if title == "" {
			fmt.Println("Please provide a title for the todo")
			return
		}
		todo.Title = title
		todo.Done = false
		todo.Created = time.Now().Format("2006-01-02 15:04:05")

		err := local.SaveTodos(todo)
		if err != nil {
			fmt.Println("Error saving todo:", err)
			return
		}
		fmt.Println("Todo added successfully")
	},
}

func init() {
	rootCmd.AddCommand(add)
}
