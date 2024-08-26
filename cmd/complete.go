package cmd

import (
	"fmt"

	"github.com/robertgouveia/TodoCLI/local"
	"github.com/spf13/cobra"
)

var completedTodo = &cobra.Command{
	Use:   "completed",
	Short: "Completed Todos",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := local.LoadTodos()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, todo := range todos {
			if todo.Done {
				fmt.Printf("✅ %s\n", todo.Title)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completedTodo)
}
