package cmd

import (
	"fmt"
	"github.com/robertgouveia/TodoCLI/local"
	"github.com/spf13/cobra"
)

var title string
var done bool

var add = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		var todo local.Todo

		if title == "" {
			fmt.Println("Error: Missing title flag")
			return
		}
		todo.Title = title
		todo.Done = done

		err := local.SaveTodos(todo)
		if err != nil {
			fmt.Println("Error saving todo:", err)
			return
		}
		fmt.Println("Todo added successfully")
	},
}

func init() {
	add.Flags().StringVarP(&title, "title", "t", "", "Title of the todo")
	add.Flags().BoolVarP(&done, "done", "d", false, "Status of the todo")
	rootCmd.AddCommand(add)
}
