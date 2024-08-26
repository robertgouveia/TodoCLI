package cmd

import (
	"fmt"

	"github.com/robertgouveia/TodoCLI/local"
	"github.com/spf13/cobra"
)

var index int

var removeTodo = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if index == 0 {
			fmt.Println("Error: Indexing starts at 1")
			return
		}

		confirm := ""
		fmt.Println("Are you sure you want to remove this todo? (y/n)")
		fmt.Scanln(&confirm)
		if confirm != "y" {
			fmt.Println("Operation Cancelled")
			return
		}

		err := local.RemoveTodos(index)

		if err != nil {
			fmt.Println("Error removing todo:", err)
			return
		}
		fmt.Println("Todo removed successfully")
	},
}

func init() {
	removeTodo.Flags().IntVarP(&index, "index", "i", 0, "Index to Remove")
	rootCmd.AddCommand(removeTodo)
}
