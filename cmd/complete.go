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

		tbl := returnTable()
		completed := 0
		for i, todo := range todos {
			if todo.Done {
				completed++
				tbl.AddRow(i, "✔ "+todo.Title, todo.Created)
			}
		}
		if completed == 0 {
			fmt.Println("No completed todos found")
			return
		}
		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(completedTodo)
}
