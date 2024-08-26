package cmd

import (
	"fmt"

	"github.com/robertgouveia/TodoCLI/local"
	"github.com/spf13/cobra"
)

var num = 0

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

		if num == 0 {
			for i, todo := range todos {
				if todo.Done {
					tbl.AddRow(i, "✔ "+todo.Title, todo.Created)
				} else {
					tbl.AddRow(i, "❌"+todo.Title, todo.Created)
				}
			}

			tbl.Print()
		} else {
			if num > len(todos) {
				fmt.Println("Error: Not enough todos")
				return
			}

			for i := 1; i <= num; i++ {
				fmt.Println(todos[i-1].Title)
			}
			tbl.Print()
		}
	},
}

func init() {
	todoList.Flags().IntVarP(&num, "num", "n", 0, "Number of Todos")

	rootCmd.AddCommand(todoList)
}
