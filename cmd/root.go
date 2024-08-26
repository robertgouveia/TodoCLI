package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A Simple Todo CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello World")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
