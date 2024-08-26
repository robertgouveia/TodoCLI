package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// defining a flag
var shout bool

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A Simple Todo CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Todo CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
