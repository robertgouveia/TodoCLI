package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints hello world",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
