package cmd

import (
	"fmt"

	"github.com/robertgouveia/TodoCLI/local"
	"github.com/spf13/cobra"
)

var all bool

var todoList = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := local.LoadTodos()
		if err != nil {
			fmt.Println("Error loading todos:", err)
			return
		}

		tbl := returnTable()

		if len(todos) == 0 {
			fmt.Println("No todos found")
			return
		}

		count := 0
		for i, todo := range todos {
			if !todo.Done {
				count++
				tbl.AddRow(i, "❌"+todo.Title, todo.Created)
			} else if all {
				tbl.AddRow(i, "✔ "+todo.Title, todo.Created)
			}
		}
		if !all && count == 0 {
			fmt.Println("No incomplete todos, try running with --all flag")
			return
		}

		tbl.Print()
	},
}

func init() {
	todoList.Flags().BoolVarP(&all, "all", "a", false, "List all todos")

	rootCmd.AddCommand(todoList)
}
